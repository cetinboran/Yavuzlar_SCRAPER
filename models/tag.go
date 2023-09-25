package models

import (
	"strings"
)

func TagInit(tagName string) *Tag {
	return &Tag{name: tagName}
}

func (t *Tag) SetClasses(classes string) {
	t.class = strings.Split(classes, ",")
}

func (t *Tag) SetIds(ids string) {
	t.ids = strings.Split(ids, ",")
}
