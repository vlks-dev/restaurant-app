package restaurant

import (
	"context"
	"github.com/cvckeboy/restaurant-app/restaurant/models"
	"github.com/cvckeboy/restaurant-app/restaurant/storage"
	"log/slog"
	"time"
)

type RestaurantStore interface {
	GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error)
	GetRestaurant(ctx context.Context, id string) (models.Restaurant, error)
	CreateRestaurant(ctx context.Context, restaurant models.Restaurant) (string, error)
	UpdateRestaurant(ctx context.Context, restaurant models.Restaurant) error
	DeleteRestaurant(ctx context.Context, id string) error
}

type Service struct {
	storage *storage.Storage
	logger  *slog.Logger
}

func NewService(storage *storage.Storage, logger *slog.Logger) *Service {
	return &Service{storage, logger}
}

func (s *Service) GetAllRestaurants(ctx context.Context) ([]models.Restaurant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.storage.GetAllRestaurants(ctx)
}
