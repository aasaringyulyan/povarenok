package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Preview
// @ID getPreview
// @Accept  json
// @Produce  json
// @Param category query string false "category"
// @Param page query string false "page"
// @Success 200 {object} scraping.Preview "Preview"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} errorResponse
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
