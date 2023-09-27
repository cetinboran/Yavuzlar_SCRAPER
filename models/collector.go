package models

import (
	"log"
	"strings"

	"github.com/cetinboran/gojson/gojson"
)

// Loop Through Data
type Each func(i int, name string)

func collectorInit(table gojson.Table, config Config) *Collector {
	// Scrapper dan aldığım config dosyasını collectorlarda kullanabilmek için ekliorm.

	return &Collector{table: table, config: config}
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
	searched := tag.name

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
		f(i+1, v)
	}
}

// Saves Data to the collection table
func (c *Collector) Save() {
	c.readableData()

	if c.config.AutoSave {
		log.Fatal("Automatic Save is on. It is recommended to turn it off when using this collector's save function.")
	}

	newData := gojson.DataInit([]string{"Searched", "Findings"}, []interface{}{c.searched, c.data}, &c.table)
	c.table.Save(newData)

}
