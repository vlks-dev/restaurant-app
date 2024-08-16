package storages

import (
	"context"
	"fmt"
	"github.com/vlks-dev/restaurant-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type AdministratorStorage struct {
	client *mongo.Client
	logger *slog.Logger
}

func NewAdministratorStorage(client *mongo.Client, logger *slog.Logger) *AdministratorStorage {
	return &AdministratorStorage{client: client, logger: logger}
}

func (s *AdministratorStorage) Register(ctx context.Context, request *models.CreateAdministratorRequest) (string, error) {
	var id string

	administratorCollection := s.client.Database("administrator").Collection("administrators")

	s.logger.Info("Registering administrator in database layer", "telephone number", request.TelephoneNumber)
	insertOne, err := administratorCollection.InsertOne(ctx, &request)
	if err != nil {
		return id, err
	}
	oid, ok := insertOne.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("не удалось преобразовать InsertedID в ObjectID")
	}

	return oid.Hex(), nil
}
