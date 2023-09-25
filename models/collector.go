package models

import "fmt"

type Loop func(i int, name string) string

func collectorInit() *Collector {
	return &Collector{}
}

func (c *Collector) SetData(data string) {
	c.data = append(c.data, data)
}

func (c *Collector) GetData() []string {
	for _, v := range c.data {
		fmt.Println(v)
		fmt.Println()
	}
	return c.data
}

func (c *Collector) Each(f Loop) {

}
