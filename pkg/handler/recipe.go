package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Recipe
// @ID getRecipe
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} scraping.Recipe "Recipe"
// @Success 200 {object} scraping.Ingredients "Ingredients"
// @Success 200 {object} scraping.Step "Step"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} errorResponse
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
