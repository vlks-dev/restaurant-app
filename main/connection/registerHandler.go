package connection

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlks-dev/restaurant-app/handlers"
	"github.com/vlks-dev/restaurant-app/services"
	"github.com/vlks-dev/restaurant-app/storages"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

type Connection struct {
	client *mongo.Client
	pool   *pgxpool.Pool
}

func NewConnection(client *mongo.Client, pool *pgxpool.Pool) *Connection {
	return &Connection{
		client: client,
		pool:   pool,
	}
}

func (c *Connection) RegisterAdministratorHandler(logger *slog.Logger, engine *gin.Engine) {
	took := make(chan time.Duration, 0)
	go func(start time.Time) {
		administratorStorage := storages.NewAdministratorStorage(c.client, logger)
		administratorService := services.NewAdministratorService(administratorStorage, logger)
		administratorHandler := handlers.NewAdministratorHandler(administratorService, logger)
		administratorHandler.Register(engine)
		took <- time.Since(start)
	}(time.Now())
	logger.Debug("Registered administrator handler", "took time", took)
}

func (c *Connection) RegisterRestaurantHandler(logger *slog.Logger, engine *gin.Engine) {
	took := make(chan time.Duration, 0)
	go func(start time.Time) {
		restaurantStorage := storages.NewRestaurantStorage(c.pool, logger)
		restaurantService := services.NewService(restaurantStorage, logger)
		restaurantHandler := handlers.NewRestaurantHandler(restaurantService, logger)
		restaurantHandler.Register(engine)
		took <- time.Since(start)
	}(time.Now())
	logger.Debug("Registered restaurant handler", "took", took)
}
