package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Preview
// @Description category=1 - Свежие рецепты
// @Description category=2 - Бульоны и супы
// @Description category=6 - Горячие блюда
// @Description category=12 - Салаты
// @Description category=15 - Закуски
// @Description category=25 - Выпечка
// @Description category=30 - Десерты
// @Description category=23 - Соусы
// @ID getPreview
// @Accept  json
// @Produce  json
// @Param category query string true "category"
// @Param page query string true "page"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /recipes [get]
func (h *Handler) getPreview(ctx *gin.Context) {
	logger := h.logger.Logger

	logger.Info("Starting get category")
	category, err := getCategory(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Got category")

	logger.Info("Starting get page")
	page, err := getPage(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Got page")

	logger.Info("Starting scraping previews")
	previews, err := h.services.Scraping.GetPreview(category, page)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Finish scraping previews")

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &previews,
	})
}
