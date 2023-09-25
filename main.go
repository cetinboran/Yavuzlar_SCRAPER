package main

import (
	"fmt"

	"github.com/cetinboran/yavuzlarscraper/models"
)

func main() {
	Scraper := models.ScraperInit()

	div := models.TagInit("div")
	div.SetClasses("anan")

	Scraper.AddTag(*div)

	fmt.Println(Scraper)
}
