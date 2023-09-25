package models

type Scraper struct {
	body      []string
	Collected []Collector
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

type Collector struct {
	searched string
	data     []string
}
