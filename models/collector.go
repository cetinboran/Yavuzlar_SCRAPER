package models

import "fmt"

func collectorInit() *Collector {
	return &Collector{}
}

func (c *Collector) SetData(data string) {
	c.data = append(c.data, data)
}

func (c *Collector) GetData() {
	fmt.Println(c.data)
}
