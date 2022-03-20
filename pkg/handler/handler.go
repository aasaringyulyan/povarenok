package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"scraping/pkg/logging"
	"scraping/pkg/service"

	_ "scraping/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/recipes", h.getPreview)
	router.GET("/recipes/search", h.getSearchPreview)
	router.GET("/recipe/show", h.getRecipe)
	router.GET("/recipe/categories", h.getCategories)

	return router
}
