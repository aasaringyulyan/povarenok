package service

import "scraping"

type Scraping interface {
	GetPreview(category string, page string) ([]scraping.Preview, error)
	GetRecipe(id string) (scraping.Recipe, error)
}

type Service struct {
	Scraping
}

func NewService() *Service {
	return &Service{
		Scraping: NewScrapingService(),
	}
}
