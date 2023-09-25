package models

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/yavuzlarscraper/utitlity"
)

func ScraperInit() *Scraper {
	return &Scraper{}
}

// Burada default value'ları koyuyorum.
func (s *Scraper) setDefaultValues() {
	s.Url = ""
}

func (s *Scraper) TakeInputs(args []cla.Input) {
	s.setDefaultValues()

	for _, i2 := range args {
		if i2.Argument == "u" {
			s.Url = i2.Value
		}

		if i2.Argument == "c" {
			s.Class = i2.Value
		}

		if i2.Argument == "i" {
			s.Id = i2.Value
		}
	}

	s.ClassNIdParser()
}

func (s *Scraper) ClassNIdParser() {
	classes := utitlity.ClassNIdParser(s.Class)
	ids := utitlity.ClassNIdParser(s.Id)

	fmt.Println(classes)
	fmt.Println()
	fmt.Println(ids)
}
