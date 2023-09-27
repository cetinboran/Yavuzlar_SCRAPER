package models

import "github.com/cetinboran/gojson/gojson"

type Scraper struct {
	body      []string
	collected []Collection
	database  *gojson.Database
	config    *Config
}

type Config struct {
	AutoSave bool
}

type Tag struct {
	name      string
	id        string
	attribute []string
	class     []string
	search    *Search
}

type Search struct {
	StartReg string
	End      string
}

type Collection struct {
	searched string
	data     []string
	table    gojson.Table
	config   Config
}
