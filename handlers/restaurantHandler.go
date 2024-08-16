package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vlks-dev/restaurant-app/models"
	"github.com/vlks-dev/restaurant-app/services"
	"log/slog"
	"net/http"
)

type RestaurantHandler struct {
	service *services.RestaurantService
	logger  *slog.Logger
}

func NewRestaurantHandler(service *services.RestaurantService, logger *slog.Logger) *RestaurantHandler {
	return &RestaurantHandler{
		service: service,
		logger:  logger,
	}
}
func (h *RestaurantHandler) Register(router *gin.Engine) {
	restaurantRouter := router.Group("/api/v1/restaurant")
	restaurantRouter.GET("/", h.GetAllRestaurants)
	restaurantRouter.POST("/", h.CreateRestaurant)
}

func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	restaurants, err := h.service.GetAllRestaurants(c)
	if err != nil {
		h.logger.Error("failed to get all restaurants", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var request models.CreateRestaurantRequest

	if err := c.ShouldBind(&request); err != nil {
		h.logger.Error("failed to parse request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateRestaurant(c, &request)
	if err != nil {
		h.logger.Error("failed to create restaurant", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create restaurant", "id": id})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}
