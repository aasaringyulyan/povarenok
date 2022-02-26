package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"scraping"
)

type ScrapingService struct{}

func NewScrapingService() *ScrapingService {
	return &ScrapingService{}
}

func (s *ScrapingService) GetPreview(category string, page string) ([]scraping.Preview, error) {
	c := colly.NewCollector()

	//TODO добавить логирование
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err.Error())
	})

	//TODO добавить логирование
	c.OnResponse(func(r *colly.Response) {
		log.Println(fmt.Sprintf("Visiting: %s", r.Request.URL))
	})

	previews := make([]scraping.Preview, 0, 200)

	c.OnHTML(".item-bl", func(e *colly.HTMLElement) {
		preview := scraping.Preview{
			Id:      e.ChildAttr("div", "data-recipe"),
			Link:    e.ChildAttr("h2 a", "href"),
			Photo:   e.ChildAttr("img", "src"),
			Name:    e.ChildText("h2 a"),
			Comment: scraping.Replace(e.ChildText("article.item-bl > p")),
			Author:  scraping.Replace(e.ChildText("div.article-footer a.user-link")),
		}

		previews = append(previews, preview)
	})

	//TODO добавить логирвоание
	if category == "recipe" {
		err := c.Visit("https://www.povarenok.ru/recipes/~" + page)
		if err != nil {
			log.Println(fmt.Sprintf("err: %s", err.Error()))
		}
	} else {
		err := c.Visit("https://www.povarenok.ru/recipes/category/" + category + "/~" + page)
		if err != nil {
			log.Println(fmt.Sprintf("err: %s", err.Error()))
		}
	}

	return previews, nil
}

func (s *ScrapingService) GetRecipe(id string) (scraping.Recipe, error) {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(fmt.Sprintf("Visiting: %s", r.Request.URL))
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
			steps = append(steps, scraping.Step{
				Photo:   el.ChildAttr("a", "href"),
				Comment: scraping.Replace(el.ChildText("div p")),
			})
		})

		recipe = scraping.Recipe{
			Id:          id,
			Name:        e.ChildText("div h1"),
			Photo:       e.ChildAttr("img", "src"),
			Comment:     scraping.Replace(e.ChildText("div.article-text p")),
			Ingredients: ingredients,
			Steps:       steps,
		}
	})

	err := c.Visit("https://www.povarenok.ru/recipes/show/" + id)
	if err != nil {
		log.Printf("err : %s", err)
	}

	return recipe, nil
}
