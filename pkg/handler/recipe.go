package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRecipe(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	recipe, err := h.services.Scraping.GetRecipe(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &recipe,
	})
}
