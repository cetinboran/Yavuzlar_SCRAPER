package models

import (
	"fmt"
	"strings"
)

func SearchInit() *Search {
	return &Search{}
}

func (s *Search) setSearch(t Tag) {
	s.setStart(t)
	s.setEnd(t)
}

func (s *Search) setStart(t Tag) {
	start := fmt.Sprintf("%v", t.Name)

	if len(t.class) > 0 {
		classes := strings.Join(t.class, " ")

		start += ` class="`
		start += classes
		start += `"`
	}

	if t.id != "" {
		start += fmt.Sprintf(` id="%v"`, t.id)
	}

	s.Start = start
}

func (s *Search) setEnd(t Tag) {
	end := fmt.Sprintf("/%v", t.Name)
	s.End = end
}
