package models

import (
	"github.com/google/uuid"
	"time"
)

type Restaurant struct {
	RestaurantId   uuid.UUID `json:"restaurant_id"`
	RestaurantName string    `json:"restaurant_name"`
	Location       string    `json:"location"`
	Schedule       time.Time `json:"schedule"`
}

type CreateRestaurantRequest struct {
	RestaurantName string `json:"restaurant_name"`
	Location       string `json:"location"`
	Schedule       string `json:"schedule"`
}

type Menu struct {
	MenuId       int       `json:"menu_id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
}
