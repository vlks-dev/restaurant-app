package server

import (
	"context"
	"errors"
	"github.com/cvckeboy/restaurant-app/utils/config"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer(cfg *config.Config, logger2 *slog.Logger) {
	engine := gin.Default()

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: engine.Handler(),
		/*		WriteTimeout: cfg.Server.Timeout.Write,
				ReadTimeout:  cfg.Server.Timeout.Read,
				IdleTimeout:  cfg.Server.Timeout.Idle,*/
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger2.Error("Error starting server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger2.Info("Shutdown Server ...", "Addr:", srv.Addr)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger2.Error("Error shutdown Server", "error", err)
	}

	select {
	case <-ctx.Done():
		logger2.Info("Timeout of 5 seconds.")
	}

	logger2.Info("Server exiting")
}
