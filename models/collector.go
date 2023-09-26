package models

import (
	"strings"
)

type Loop func(i int, name string) string

func collectorInit() *Collector {
	return &Collector{}
}

func (c *Collector) setData(data string) {
	c.data = append(c.data, data)
}

func (c *Collector) setSearched(tag Tag) {
	searched := tag.Name

	if len(tag.class) > 0 {
		searched += ":" + strings.Join(tag.class, ".")
	}

	if tag.id != "" {
		searched += ": #" + tag.id
	}

	c.searched = searched
}

func (c *Collector) GetData() []string {
	return c.data
}

func (c *Collector) Each(f Loop) {

}
