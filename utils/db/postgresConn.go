package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlks-dev/restaurant-app/utils/config"
	"log/slog"
	"time"
)

func NewPostgresConn(ctx context.Context, config *config.Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.PostgreSQL.Username,
		config.Database.PostgreSQL.Password,
		config.Database.PostgreSQL.Host,
		config.Database.PostgreSQL.Port,
		config.Database.PostgreSQL.Database,
	)

	configDb, err := pgxpool.ParseConfig(url)
	if err != nil {
		logger.Error("Failed to parse postgres connection url", "url", url, "err", err)
		return nil, err
	}

	configDb.MaxConns = 100
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
