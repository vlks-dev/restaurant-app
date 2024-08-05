package models

type Restaurant struct {
	RestaurantId   int    `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
	MenuId         int    `json:"menu_id"`
	Location       string `json:"location"`
	Schedule       string `json:"schedule"`
}

type Menu struct {
	MenuId int `json:"menu_id"`
	DishId int `json:"dish_id"`
}
