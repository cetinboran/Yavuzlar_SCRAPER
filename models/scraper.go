package models

func ScraperInit() *Scraper {
	return &Scraper{tags: make([]Tag, 0)}
}

func (s *Scraper) AddTag(tags ...Tag) {
	s.tags = tags
}
