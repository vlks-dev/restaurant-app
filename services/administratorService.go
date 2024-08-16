package services

import (
	"context"
	"github.com/vlks-dev/restaurant-app/models"
	"github.com/vlks-dev/restaurant-app/storages"
	"github.com/vlks-dev/restaurant-app/utils"
	"log/slog"
)

type AdministratorStorage interface {
	Register(ctx context.Context, request *models.CreateAdministratorRequest) (string, error)
	//	Login(ctx context.Context, request *models.CreateAdministratorRequest) (*models.Administrator, error)
}

type AdministratorService struct {
	storage *storages.AdministratorStorage
	logger  *slog.Logger
}

func NewAdministratorService(storage *storages.AdministratorStorage, logger *slog.Logger) *AdministratorService {
	return &AdministratorService{
		storage: storage,
		logger:  logger,
	}
}

func (i *AdministratorService) Register(ctx context.Context, request *models.CreateAdministratorRequest) (string, error) {
	var err error
	request.Passerial, err = utils.HashPassword(request.Passerial)
	if err != nil {
		return "", err
	}
	i.logger.Debug("Registering administrator", "request", &request)
	return i.storage.Register(ctx, request)
}
