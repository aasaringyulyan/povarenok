package scraping

import "strings"

type Preview struct {
	Id      string `json:"id"`
	Link    string `json:"link"`
	Photo   string `json:"photo"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Author  string `json:"author"`
}

type Recipe struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Photo       string        `json:"photo"`
	Comment     string        `json:"comment"`
	Ingredients []Ingredients `json:"ingredients"`
	Steps       []Step        `json:"steps"`
}

type Ingredients struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Step struct {
	Photo   string `json:"photo"`
	Comment string `json:"comment"`
}

func Replace(str string) string {
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	return str
}
