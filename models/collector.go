package models

import (
	"log"
	"strings"

	"github.com/cetinboran/gojson/gojson"
)

// Loop Through Data
type Each func(i int, name string)

func collectionInit(table gojson.Table, config Config) *Collection {
	// Scrapper dan aldığım config dosyasını Collectionlarda kullanabilmek için ekliorm.

	return &Collection{table: table, config: config}
}

func (c *Collection) setData(data string) {
	c.data = append(c.data, data)
}

// Changes Collection data with user readable data
func (c *Collection) readableData() {
	var newDataArr []string

	for _, v := range c.data {
		for _, v2 := range strings.Split(v, "\n") {
			newDataArr = append(newDataArr, v2)
		}
	}

	c.data = newDataArr
}

func (c *Collection) setSearched(tag Tag) {
	searched := tag.name

	if len(tag.class) > 0 {
		searched += ":" + strings.Join(tag.class, ".")
	}

	if tag.id != "" {
		searched += ": #" + tag.id
	}

	c.searched = searched
}

func (c *Collection) Get() []string {
	return c.data
}

func (c *Collection) Each(f Each) {
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
func (c *Collection) Save() {
	c.readableData()

	if c.config.AutoSave {
		log.Fatal("Automatic Save is on. It is recommended to turn it off when using this Collection's save function.")
	}

	newData := gojson.DataInit([]string{"Searched", "Findings"}, []interface{}{c.searched, c.data}, &c.table)
	c.table.Save(newData)

}
