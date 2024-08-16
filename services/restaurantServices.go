package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/vlks-dev/restaurant-app/models"
	"github.com/vlks-dev/restaurant-app/storages"
	"log/slog"
	"time"
)

type RestaurantStore interface {
	GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error)
	CreateRestaurant(ctx context.Context, req *models.CreateRestaurantRequest) (uuid.UUID, error)
	/*	GetRestaurant(ctx context.Context, id string) (models.Restaurant, error)

		UpdateRestaurant(ctx context.Context, restaurant models.Restaurant) error
		DeleteRestaurant(ctx context.Context, id string) error*/
}

type RestaurantService struct {
	storage *storages.RestaurantStorage
	logger  *slog.Logger
}

func NewService(storage *storages.RestaurantStorage, logger *slog.Logger) *RestaurantService {
	return &RestaurantService{storage, logger}
}

func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *models.CreateRestaurantRequest) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.storage.CreateRestaurant(ctx, req)
}

func (s *RestaurantService) GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.storage.GetAllRestaurants(ctx)
}
