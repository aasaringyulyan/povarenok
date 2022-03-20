package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"scraping"
	"strings"
)

func scrapingPreview(url string, page string, logger *logrus.Logger) ([]scraping.Preview, error) {
	c := colly.NewCollector()

	var requestURL string

	previews := make([]scraping.Preview, 0, 200)

	c.OnError(func(_ *colly.Response, err error) {
		logger.Infof("Error: %s", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		logger.Infof(fmt.Sprintf("Visiting: %s", r.Request.URL))
		requestURL = r.Request.URL.String()
	})

	c.OnHTML(".item-bl", func(e *colly.HTMLElement) {
		if requestURL != url && page != "1" {
			return
		}

		preview := scraping.Preview{
			Id:      e.ChildAttr("div", "data-recipe"),
			Photo:   e.ChildAttr("img", "src"),
			Name:    e.ChildText("h2 a"),
			Comment: scraping.Replace(e.ChildText("article.item-bl > p")),
			Author:  scraping.Replace(e.ChildText("div.article-footer a.user-link")),
		}

		previews = append(previews, preview)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return previews, nil
}

func scrapingRecipe(url string, id string, logger *logrus.Logger) (scraping.Recipe, error) {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		logger.Infof("Error: %s", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		logger.Infof(fmt.Sprintf("Visiting: %s", r.Request.URL))
	})

	var recipe scraping.Recipe

	c.OnHTML(".item-bl", func(e *colly.HTMLElement) {
		ingredients := make([]scraping.Ingredients, 0)
		steps := make([]scraping.Step, 0)

		e.ForEach("div.ingredients-bl ul li", func(_ int, el *colly.HTMLElement) {
			ingredients = append(ingredients, scraping.Ingredients{
				Name:  scraping.Replace(el.ChildText("a")),
				Value: scraping.Replace(el.ChildText("span > span")),
			})
		})

		e.ForEach("div.cooking-bl", func(_ int, el *colly.HTMLElement) {
			points := strings.Split(el.ChildText("div p"), "\n")

			steps = append(steps, scraping.Step{
				Photo:   el.ChildAttr("a", "href"),
				Comment: points,
			})
		})

		recipe = scraping.Recipe{
			Id:          id,
			Name:        e.ChildText("div h1"),
			Photo:       e.ChildAttr("img[itemprop=image]", "src"),
			Comment:     scraping.Replace(e.ChildText("div.article-text p")),
			Ingredients: ingredients,
			Steps:       steps,
		}
	})

	err := c.Visit(url)
	if err != nil {
		return scraping.Recipe{}, err
	}

	return recipe, nil
}
