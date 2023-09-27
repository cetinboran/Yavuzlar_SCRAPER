package models

import (
	"strings"
)

type Each func(i int, name string)

func collectorInit() *Collector {
	return &Collector{}
}

func (c *Collector) setData(data string) {
	c.data = append(c.data, data)
}

// Changes collector data with user readable data
func (c *Collector) readableData() {
	var newDataArr []string

	for _, v := range c.data {
		for _, v2 := range strings.Split(v, "\n") {
			newDataArr = append(newDataArr, v2)
		}
	}

	c.data = newDataArr
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

func (c *Collector) Each(f Each) {
	var clearData []string

	for _, v := range c.data {
		for _, v2 := range strings.Split(v, "\n") {
			clearData = append(clearData, v2)
		}
	}

	for i, v := range clearData {
		f(i, v)
	}
}
