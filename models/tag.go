package models

import (
	"strings"
)

func TagInit(tagName string) *Tag {
	return &Tag{Name: tagName, Search: SearchInit()}
}

func (t *Tag) SetClasses(classes string) {
	t.class = strings.Split(classes, ",")
}

func (t *Tag) SetId(id string) {
	t.id = id
}

func (t *Tag) setSearch() {
	t.Search.setSearch(*t)
}
