package handler

import (
	"github.com/gin-gonic/gin"
	"scraping/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/recipes", h.getPreview)
	router.GET("/recipes/show", h.GetRecipe)

	return router
}
