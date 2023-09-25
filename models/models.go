package models

type Scraper struct {
	tags []Tag
}

type Tag struct {
	Name   string
	id     string
	class  []string
	Search *Search
}

type Search struct {
	Start string
	End   string
}
