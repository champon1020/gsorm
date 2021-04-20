package internal

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// SnakeCase returns str as snake case.
func SnakeCase(str string) (snake string) {
	str = strings.Split(str, "#")[0]
	snake = matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ToLower(snake)
	return
}

// ToStringOpt is the option of ToString.
type ToStringOpt struct {
	Quotes     bool
	TimeFormat string
}

// ToString converts the type of value to string.
// If quotes is true, it attaches single quotes to returned value.
// Default conversion format is as follows:
//  str (string)                            -> "str" (If quotes == true, "'str'")
//  0 (int, intN)                           -> "0"
//  0 (uint, uintN)                         -> "0"
//  1.0 (floatN)                            -> "1.00000"
//  true (bool)                             -> "1" (If false, "0")
//  2006-01-02T15:04:05Z00:00 (time.Time)   -> "2006-01-02 15:04:05"
//  nil                                     -> "nil"
func ToString(v interface{}, opt *ToStringOpt) string {
	if v == nil {
		return "nil"
	}

	if opt == nil {
		opt = &ToStringOpt{Quotes: true, TimeFormat: "2006-01-02 15:04:05"}
	}

	switch v := v.(type) {
	case string:
		if opt.Quotes {
			return fmt.Sprintf("'%s'", v)
		}
		return v
	case int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%v", v), "0"), ".")
	case bool:
		if v {
			return "1"
		}
		return "0"
	case time.Time:
		t := v.Format(opt.TimeFormat)
		if opt.Quotes {
			return fmt.Sprintf("'%s'", t)
		}
		return t
	}

	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice || typ == reflect.Array {
		var s string
		vals := reflect.ValueOf(v)
		for i := 0; i < vals.Len(); i++ {
			if i != 0 {
				s += ", "
			}
			s += ToString(vals.Index(i).Interface(), opt)
		}
		return s
	}

	return reflect.TypeOf(v).String()
}

// ColumnsAndFields generates map of column index and field index.
func ColumnsAndFields(cols []string, modelTyp reflect.Type) map[int]int {
	candf := make(map[int]int)
	for i, c := range cols {
		for j := 0; j < modelTyp.NumField(); j++ {
			if c != ColumnName(modelTyp.Field(j)) {
				continue
			}
			candf[i] = j
		}
	}
	return candf
}

// ColumnName returns column name with field tag.
func ColumnName(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return SnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}
