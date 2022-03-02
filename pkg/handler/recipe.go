package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Recipe
// @Tags recipe
// @Description Get recipe
// @ID getRecipe
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /recipe/show [get]
func (h *Handler) getRecipe(ctx *gin.Context) {
	logger := h.logger.Logger

	logger.Info("Starting get recipe id")
	id, err := getId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Got recipe id")

	logger.Infof("Starting scraping recipe (id=%s)", id)
	recipe, err := h.services.Scraping.GetRecipe(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("Finish scraping recipe")

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &recipe,
	})
}
