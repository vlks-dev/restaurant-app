package db

import (
	"context"
	"fmt"
	"github.com/vlks-dev/restaurant-app/utils/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"time"
)

func NewMongoConn(ctx context.Context, config *config.Config, logger *slog.Logger) (*mongo.Client, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	uri := fmt.Sprintf(
		"mongodb://%v:%v",
		config.Database.MongoDB.Host,
		config.Database.MongoDB.Port,
	)
	opts := options.Client().ApplyURI(uri)
	logger.Debug(
		"connecting to mongodb",
		"connectionUrl", uri,
	)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Error("failed to connect to mongodb", "error", err)
		return nil, err
	}

	return client, nil
}
