package models

import (
	"regexp"
	"strings"
)

func ScraperInit() *Scraper {
	return &Scraper{}
}

func (s *Scraper) SetBody(body []string) {
	s.body = body
}

func (s *Scraper) getText(start, end int) string {
	body := strings.Join(s.body[start:end], "")

	re := regexp.MustCompile(">([^<]+)")

	matches := re.FindAllStringSubmatch(body, -1)

	var data []string
	for _, match := range matches {
		match[1] = strings.TrimSpace(match[1])
		data = append(data, match[1])
	}

	return strings.Join(data, " ")
}

func (s *Scraper) Find(tag Tag) *Collector {
	newCollector := collectorInit()

	tag.Search.setSearch(tag)

	startTag := tag.Search.Start
	endTag := tag.Search.End

	var startIndex, endIndex int

	var i int
	for i < len(s.body) {

		if strings.Contains(s.body[i], startTag) {
			startIndex = i
			for {
				i++

				if strings.Contains(s.body[i], endTag) {
					endIndex = i
					break
				}
			}

			// Eğer döngüden çıktyısa içeriden texti çıkar
			data := s.getText(startIndex-1, endIndex+1)
			newCollector.SetData(data)
		}

		i++
	}

	s.Collected = append(s.Collected, *newCollector)
	return &s.Collected[len(s.Collected)-1]
}
