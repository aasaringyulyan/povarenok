package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getPreview(ctx *gin.Context) {
	category, err := getCategory(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	page, err := getPage(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	previews, err := h.services.Scraping.GetPreview(category, page)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &previews,
	})
}
