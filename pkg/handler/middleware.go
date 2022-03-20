package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/charmap"
	"strconv"
)

const (
	pageCtx        = "page"
	categoryCtx    = "category"
	idCtx          = "id"
	searchInputCtx = "searchInput"
)

// this is meant to be constant! Please don't mutate it!
var availableCategory = []string{"1", "2", "6", "12", "15", "25", "30", "23"}

func getPage(ctx *gin.Context) (string, error) {
	page := ctx.DefaultQuery(pageCtx, "1")

	_, ok := strconv.ParseUint(page, 10, 32)
	if ok != nil {
		return "", errors.New("page is of invalid type")
	}

	return page, nil
}

func getCategory(ctx *gin.Context) (string, error) {
	category := ctx.DefaultQuery(categoryCtx, "1")
	for _, n := range availableCategory {
		if category == n {
			return category, nil
		}
	}

	return "", errors.New("category not found")
}

func getSearchInput(ctx *gin.Context) (string, error) {
	searchInput := ctx.Query(searchInputCtx)

	encoder := charmap.Windows1251.NewEncoder()
	searchInput, err := encoder.String(searchInput)
	if err != nil {
		return "", errors.New("searchInput is of invalid type")
	}

	return searchInput, nil
}

func getId(ctx *gin.Context) (string, error) {
	id := ctx.Query(idCtx)

	intId, ok := strconv.Atoi(id)
	if ok != nil || intId < 0 {
		return "", errors.New("id is of invalid type")
	}

	return id, nil
}
