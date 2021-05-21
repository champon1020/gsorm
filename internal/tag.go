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
}

// Lookup returns tag exists or not.
func (t *Tag) Lookup(tag string) bool {
	switch tag {
	case "col":
		return t.Column != ""
	case "typ":
		return t.Type != ""
	case "notnull":
		return t.NotNull
	case "default":
		return t.Default != ""
	case "pk":
		return t.PK != ""
	case "fk":
		return t.FK != "" && t.Ref != ""
	case "uc":
		return t.UC != ""
	}
	return false
}

// ExtractTags extracts the struct field tags from struct.
func ExtractTags(t reflect.Type) []*Tag {
	tags := make([]*Tag, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		tags[i] = ExtractTag(t.Field(i))
	}
	return tags
}

// ExtractTag extracts the struct field tag.
func ExtractTag(f reflect.StructField) *Tag {
	t := &Tag{}
	jsonTag := f.Tag.Get("json")
	if jsonTag != "" {
		t.Column = jsonTag
	}

	tags := strings.Split(f.Tag.Get("mgorm"), ",")
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
		}
	}
	return t
}

/*
// timeFormat returns layout of date.
func timeFormat(layout string) string {
	switch layout {
	case "time.ANSIC":
		return time.ANSIC
	case "time.UnixDate":
		return time.UnixDate
	case "time.RubyDate":
		return time.RubyDate
	case "time.RFC822":
		return time.RFC822
	case "time.RFC822Z":
		return time.RFC822Z
	case "time.RFC850":
		return time.RFC850
	case "time.RFC1123":
		return time.RFC1123
	case "time.RFC1123Z":
		return time.RFC1123Z
	case "time.RFC3339":
		return time.RFC3339
	case "time.RFC3339Nano":
		return time.RFC3339Nano
	case "time.Kitchen":
		return time.Kitchen
	case "time.Stamp":
		return time.Stamp
	case "time.StampMilli":
		return time.StampMilli
	case "time.StampMicro":
		return time.StampMicro
	case "time.StampNano":
		return time.StampNano
	}
	return layout
}
*/
