package models

import (
	"fmt"
	"regexp"
	"strings"
)

func SearchInit() *Search {
	return &Search{}
}

// Sets all the things we need.
func (s *Search) setSearch(t Tag) {
	s.createRegex(t)
	s.setEnd(t)
}

func (s *Search) createRegex(t Tag) {
	// <div[^>]*class="selam"[^>]*id="3"[^>]*> => DORU BU
	regex := "<"

	// Eğer tag ismi girilmediyse sadece girilenlere göre arasın.
	if t.name != "" {
		regex += t.name
	}

	if len(t.attribute) > 0 {

		for _, v := range t.attribute {
			regex += `[^>]*` + v
		}
	}

	if len(t.class) > 0 {
		classes := strings.Join(t.class, " ")

		regex += `[^>]*class="`
		regex += classes
		regex += `"`
	}

	if t.id != "" {
		regex += `[^>]*`
		regex += fmt.Sprintf(`id="%v"`, t.id)
	}

	regex += `[^>]*>`

	s.StartReg = regex
}

func (s *Search) getAttributeRegex(attribute string) string {
	// <[^>]*href="[^>]*"[^>]*>
	regex := `<[^>]*` + attribute + `="[^>]*"[^>]*>`

	return regex
}

func (s *Search) RegexCheck(t Tag, data string) bool {
	// Regex'i işliyorum
	reg := regexp.MustCompile(s.StartReg)

	// gelen Data'da regex var ise true yok ise false dönüyor.
	return reg.MatchString(data)
}

func (s *Search) setEnd(t Tag) {
	end := fmt.Sprintf("/%v", t.name)
	s.End = end
}
