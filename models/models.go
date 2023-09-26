package models

type Scraper struct {
	head      []string
	body      []string
	Collected []Collector
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
	Start    string
	End      string
}

type Collector struct {
	searched string
	data     []string
}
