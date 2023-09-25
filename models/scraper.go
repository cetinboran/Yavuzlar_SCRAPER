package models

import (
	"fmt"
	"log"
	"strings"

	cla "github.com/cetinboran/goarg/CLA"
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
	// Alttaki formatta göre gelicek.
	// div:title,red; artical;  a:link
	// başında tag ismi olanlar sadece ona bağlı olmayanlar ise global bütün gelen -t eklenmiş gibi olucak.
	// istediğim format ise bir map[string][string]

	class := make(mapClassId)

	classGroup := strings.Split(s.Class, ";")

	var globalClasses string
	for _, v := range classGroup {
		if strings.Count(v, ":") != 1 && strings.Count(v, ":") != 0 {
			log.Fatal("Please Enter spesifc format.")
		}

		if strings.Count(v, ":") == 0 {
			globalClasses += v + ","
		} else {
			tagAndClass := strings.Split(v, ":")

			tag := tagAndClass[0]
			classes := tagAndClass[1]

			class[tag] = classes + ","
		}
	}

	// Virgülden kurtuluyorum.
	if len(globalClasses) > 0 {
		globalClasses = globalClasses[:len(globalClasses)-1]
	}

	globalClassesArr := strings.Split(globalClasses, ",")

	for k := range class {
		for _, v := range globalClassesArr {
			class[k] += v + ","
		}

		if len(globalClasses) > 0 {
			class[k] = class[k][:len(class[k])-1]
		} else {
			class[k] = class[k][:len(class[k])-2]
		}
	}

	for k, v := range class {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)
	}
}
