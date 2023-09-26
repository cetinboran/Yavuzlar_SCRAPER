package models

import "github.com/cetinboran/gojson/gojson"

type Scraper struct {
	body      []string
	Collected []Collector
	database  *gojson.Database
	config    *Config
}

type Config struct {
	AutoSave bool
}

type Tag struct {
	Name   string
	id     string
	class  []string
	Search *Search
}

type Search struct {
	StartReg string
	End      string
}

type Collector struct {
	searched string
	data     []string
}
