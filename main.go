package main

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/yavuzlarscraper/handleinput"
)

func main() {
	Setup := cla.Init()

	Setup.SetUsage("Yavuzlar Scraper", "This is web scraper.", []string{})

	Setup.AddOptionTitle("INPUT OPTION")
	Setup.AddOption("-u", false, "Enter the url")

	Setup.AddOptionTitle("SEARCH OPTION")
	Setup.AddOption("-t", false, "Set tag. a;href,div")
	Setup.AddOption("-c", false, "Search with class.")
	Setup.AddOption("-i", false, "Search with id.")

	Setup.AutomaticUsage()

	args, errors := Setup.Start()

	handleinput.Handle(args, errors)

}
