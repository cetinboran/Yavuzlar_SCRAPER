package models

type mapClassId map[string]string

type Scraper struct {
	Url   string
	Class string
	Id    string
}
