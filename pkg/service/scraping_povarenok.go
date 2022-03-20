package service

import (
	"scraping"
	"scraping/pkg/logging"
)

const (
	visitShow     = "https://www.povarenok.ru/recipes/show/"
	visitCategory = "https://www.povarenok.ru/recipes/category/"
	visitRecipe   = "https://www.povarenok.ru/recipes/~"
	visitSearch   = "https://www.povarenok.ru/recipes/search/~"
)

type ScrapingService struct {
	logger *logging.Logger
}

func NewScrapingService(logger *logging.Logger) *ScrapingService {
	return &ScrapingService{
		logger: logger,
	}
}

func (s *ScrapingService) GetPreview(category string, page string) ([]scraping.Preview, error) {
	logger := s.logger.Logger

	url := visitCategory + category + "/~" + page + "/"

	if category == "1" {
		url = visitRecipe + page + "/"
	}

	previews, err := scrapingPreview(url, page, logger)
	if err != nil {
		return nil, err
	}

	return previews, nil
}

func (s *ScrapingService) GetSearchPreview(searchInput string, page string) ([]scraping.Preview, error) {
	logger := s.logger.Logger

	url := visitSearch + page + "/?name=" + searchInput

	previews, err := scrapingPreview(url, page, logger)
	if err != nil {
		return nil, err
	}

	return previews, nil
}

func (s *ScrapingService) GetRecipe(id string) (scraping.Recipe, error) {
	logger := s.logger.Logger

	url := visitShow + id

	recipe, err := scrapingRecipe(url, id, logger)
	if err != nil {
		return scraping.Recipe{}, err
	}

	return recipe, nil
}
