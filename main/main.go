package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vlks-dev/restaurant-app/main/connection"
	"github.com/vlks-dev/restaurant-app/utils/config"
	"github.com/vlks-dev/restaurant-app/utils/db"
	"github.com/vlks-dev/restaurant-app/utils/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.NewConfig("config.yaml")
	handler := logger.MyLogger(cfg)
	logger := slog.New(handler)

	engine := gin.Default()
	ctx := context.Background()

	pool, err := db.NewPostgresConn(ctx, cfg, logger)
	if err != nil {
		logger.Error("Failed to connect to postgres database", "error", err)
		return
	}
	defer pool.Close()

	client, err := db.NewMongoConn(ctx, cfg, logger)
	if err != nil {
		logger.Error("Failed to connect to mongo database", "error", err)
		return
	}
	defer client.Disconnect(ctx)

	newConnection := connection.NewConnection(client, pool)

	newConnection.RegisterRestaurantHandler(logger, engine)
	newConnection.RegisterAdministratorHandler(logger, engine)

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("Failed to ping mongo database", "error", err)
		return
	}

	StartServer(cfg, logger, engine)
}

func StartServer(cfg *config.Config, logger *slog.Logger, engine *gin.Engine) {
	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: engine.Handler(),

		/*		WriteTimeout: cfg.Server.Timeout.Write,
				ReadTimeout:  cfg.Server.Timeout.Read,
				IdleTimeout:  cfg.Server.Timeout.Idle,*/
	}

	logger.Info("Config loaded, logger set", "program level", cfg.Server.Level, "Addr:", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Error starting server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Shutdown Server ...", "Addr:", srv.Addr)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Error shutdown Server", "error", err)
	}

	select {
	case <-ctx.Done():
		logger.Info("Timeout of 5 seconds.")
	}

	logger.Info("Server exiting")
}
