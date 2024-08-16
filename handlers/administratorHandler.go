package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vlks-dev/restaurant-app/models"
	"github.com/vlks-dev/restaurant-app/services"
	"log/slog"
)

type AdministratorHandler struct {
	service *services.AdministratorService
	logger  *slog.Logger
}

func NewAdministratorHandler(service *services.AdministratorService, logger *slog.Logger) *AdministratorHandler {
	return &AdministratorHandler{
		service: service,
		logger:  logger,
	}
}

func (h *AdministratorHandler) Register(router *gin.Engine) {
	adminRouter := router.Group("/api/v1/administrator")
	adminRouter.POST("/register", h.RegisterAdmin)
}

func (h *AdministratorHandler) RegisterAdmin(c *gin.Context) {
	var request models.CreateAdministratorRequest

	if err := c.ShouldBind(&request); err != nil {
		h.logger.Error("failed to parse request", "err", err)
		c.JSON(400, gin.H{"message": err.Error()})
	}

	id, err := h.service.Register(c, &request)
	if err != nil {
		h.logger.Error("failed to register administrator", "err", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	h.logger.Info("registered administrator", "id", id)
	c.JSON(200, gin.H{"id": id})
}
