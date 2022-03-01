package handler

import (
	"github.com/gin-gonic/gin"
	"scraping/pkg/logging"
	"scraping/pkg/service"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(logger *logging.Logger, services *service.Service) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/recipes", h.getPreview)
	router.GET("/recipes/show", h.GetRecipe)

	return router
}
