package models

import (
	"github.com/google/uuid"
)

type Administrator struct {
	AdministratorID uuid.UUID `json:"administrator_id"`
	TelephoneNumber int       `json:"telephone_number"`
	Passerial       string    `json:"-"`
	IsVerified      bool      `json:"is_verified"`
}

type CreateAdministratorRequest struct {
	TelephoneNumber int    `bson:"telephone_number"`
	Passerial       string `bson:"passerial"`
}
