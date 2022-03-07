package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scraping"
)

// @Tags Categories
// @ID getCategories
// @Accept  json
// @Produce  json
// @Success 200 {object} scraping.Category "Category"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} errorResponse
// @Router /recipe/categories [get]
func (h *Handler) getCategories(ctx *gin.Context) {
	logger := h.logger.Logger

	logger.Info("Starting get categories")
	categories := make([]scraping.Category, 8)

	categories = []scraping.Category{
		{
			Id:   "1",
			Name: "Свежие рецепты",
		},
		{
			Id:   "2",
			Name: "Бульоны и супы",
		},
		{
			Id:   "6",
			Name: "Горячие блюда",
		},
		{
			Id:   "12",
			Name: "Салаты",
		},
		{
			Id:   "15",
			Name: "Закуски",
		},
		{
			Id:   "25",
			Name: "Выпечка",
		},
		{
			Id:   "30",
			Name: "Десерты",
		},
		{
			Id:   "23",
			Name: "Соусы",
		},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &categories,
	})
}
