package scraper

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/yavuzlarscraper/models"
)

func Start(args []cla.Input, errors cla.GoargErrors) {
	Scraper := models.ScraperInit()
	Scraper.TakeInputs(args)

	// fmt.Println(Scraper)
}
