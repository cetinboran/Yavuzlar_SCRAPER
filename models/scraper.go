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

	tagCount := 0

	var i int
	for i < len(s.body) {
		// Eğer startTag'ı içeriyorsa satır o zaman içeri giriyorum.
		if strings.Contains(s.body[i], startTag) {
			// TagCount burada bulduğum tag sayısı
			tagCount++

			// startIndex buraya ilk girdiğim kısım oluyor.
			startIndex := i

			// Bu fonksiyon ile endIndexsi buluyorum.
			endIndex := s.findEndIndex(tag.Search.End, i, tagCount)

			// Ustteki fonksiyonda i yi de yolladığım için 0 yapıyorum.
			tagCount = 0

			// data'yı collector'a ekliyorum.
			var data string
			if endIndex == -1 {
				// Eğer last ındex yok ise belki end tagı olmayan input felandır
				// o zaman tagın ilk görüldüğü satırı ekliyorum.
				data = s.getText(startIndex-1, startIndex+1)
			} else {
				data = s.getText(startIndex-1, endIndex+1)
			}

			newCollector.SetData(data)
		}

		i++
	}

	s.Collected = append(s.Collected, *newCollector)
	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) findEndIndex(endTag string, start, c int) int {
	tagCount := 0
	i := start
	for i < len(s.body) {
		if strings.Contains(s.body[i], endTag) {
			// endIndex = i
			tagCount++
		}

		if c == tagCount {
			return i
		}
		i++
	}

	return -1
}
