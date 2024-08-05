package db

import (
	"context"
	"fmt"
	"github.com/cvckeboy/restaurant-app/utils/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

func NewPostgresConn(ctx context.Context, config *config.Config, logger *slog.Logger) (*pgxpool.Pool, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Database.Username, config.Database.Password,
		config.Database.Host, config.Database.Port, config.Database.Database)

	configDb, err := pgxpool.ParseConfig(url)
	if err != nil {
		logger.Error("Failed to parse postgres connection url", "url", url, "err", err)
		return nil, err
	}

	configDb.MaxConns = 100
	configDb.ConnConfig.RuntimeParams = map[string]string{
		"parseTime": "true",
	}
	configDb.MinConns = 1
	configDb.MaxConnIdleTime = 55 * time.Second
	configDb.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, configDb)
	if err != nil {
		logger.Error("Failed to create postgres connection pool", "url", url, "err", err)
		return nil, err
	}

	return pool, nil
}
