package handleinput

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/yavuzlarscraper/scraper"
)

func Handle(args []cla.Input, errors cla.GoargErrors) {
	modeName := args[0].ModeName

	switch modeName {
	case "Main":
		scraper.Start(args, errors)
	}
}
