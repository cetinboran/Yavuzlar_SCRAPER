package models

type Scraper struct {
	tags []Tag
}

type Tag struct {
	name  string
	ids   []string
	class []string
}
