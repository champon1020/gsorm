package internal

import (
	"reflect"
	"strings"
)

// Tag stores the field tag contents.
type Tag struct {
	Column  string
	Type    string
	NotNull bool
	Default string
	PK      string
	FK      string
	Ref     string
	UC      string
	Layout  string
}

// ExtractTag extracts the field tag.
func ExtractTag(f reflect.StructField) *Tag {
	t := &Tag{}
	jsonTag := f.Tag.Get("json")
	if jsonTag != "" {
		t.Column = jsonTag
	}

	tags := strings.Split(f.Tag.Get("mgorm"), " ")
	for _, v := range tags {
		if !strings.Contains(v, "=") {
			t.Column = v
			continue
		}

		eq := strings.Split(v, "=")
		switch eq[0] {
		case "typ":
			t.Type = eq[1]
		case "notnull":
			if eq[1] == "t" {
				t.NotNull = true
			}
		case "default":
			t.Default = eq[1]
		case "pk":
			t.PK = eq[1]
		case "fk":
			fk := strings.Split(eq[1], ":")
			t.FK = fk[0]
			if len(fk) == 2 {
				t.Ref = fk[1]
			}
		case "uc":
			t.UC = eq[1]
		case "layout":
			t.Layout = eq[1]
		}
	}
	return t
}
