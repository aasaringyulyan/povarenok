package service

import (
	"scraping"
	"scraping/pkg/logging"
)

type Scraping interface {
	GetPreview(category string, page string) ([]scraping.Preview, error)
	GetRecipe(id string) (scraping.Recipe, error)
	GetSearchPreview(name string, page string) ([]scraping.Preview, error)
}

type Service struct {
	Scraping
}

func NewService(logger *logging.Logger) *Service {
	return &Service{
		Scraping: NewScrapingService(logger),
	}
}
