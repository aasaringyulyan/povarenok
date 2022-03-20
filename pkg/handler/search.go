package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Preview
// @ID getSearchPreview
// @Accept  json
// @Produce  json
// @Param searchInput query string true "searchInput"
// @Param page query string false "page"
// @Success 200 {object} scraping.Preview "Preview"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} errorResponse
// @Router /recipes/search [get]
func (h *Handler) getSearchPreview(ctx *gin.Context) {
	logger := h.logger.Logger

	logger.Info("Starting get value")
	searchInput, err := getSearchInput(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Got name")

	logger.Info("Starting get page")
	page, err := getPage(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Got page")

	logger.Info("Starting scraping previews")
	previews, err := h.services.Scraping.GetSearchPreview(searchInput, page)
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
