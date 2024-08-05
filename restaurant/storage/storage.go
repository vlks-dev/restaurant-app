package storage

import (
	"context"
	"github.com/cvckeboy/restaurant-app/restaurant/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Storage struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewStorage(pool *pgxpool.Pool, logger *slog.Logger) *Storage {
	return &Storage{pool, logger}
}

func (s *Storage) GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error) {
	s.logger.Info("Getting list of all restaurants")

	q := "SELECT restaurant_name, location, schedule FROM restaurant"
	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant
	for rows.Next() {
		var req models.Restaurant
		if err := rows.Scan(&req.RestaurantName, &req.Location, &req.Schedule); err != nil {
			s.logger.Error("Error getting restaurants", err)
			return nil, err
		}
		restaurants = append(restaurants, req)
	}
	s.logger.Info("Get list of restaurant", "len", len(restaurants))
	return restaurants, nil
}
