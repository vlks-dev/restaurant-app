package storages

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlks-dev/restaurant-app/models"
	"log/slog"
)

type RestaurantStorage struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewRestaurantStorage(pool *pgxpool.Pool, logger *slog.Logger) *RestaurantStorage {
	return &RestaurantStorage{pool, logger}
}

func (s *RestaurantStorage) CreateRestaurant(ctx context.Context, req *models.CreateRestaurantRequest) (uuid.UUID, error) {
	s.logger.Info("Inserting in restaurant table", "restaurant name", req.RestaurantName)
	var id uuid.UUID

	q := `
INSERT INTO public.restaurant (restaurant_name, location, schedule)
VALUES ( $1, $2, $3 )
RETURNING restaurant_id;`

	err := s.pool.QueryRow(ctx, q, req.RestaurantName, req.Location, req.Schedule).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	s.logger.Info("Inserted restaurant", "id", id)
	return id, nil
}

func (s *RestaurantStorage) GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error) {
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

	if len(restaurants) == 0 {
		s.logger.Error("No restaurants found")
		newErr := errors.New("no restaurants found")
		return nil, newErr
	}

	return restaurants, nil
}
