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

// ToString convert the value of interface to string.
// If quotes is true, attache double quotes to string value.
func ToString(v interface{}, quotes bool) string {
	if v == nil {
		return "nil"
	}

	switch v := v.(type) {
	case string:
		if quotes {
			return fmt.Sprintf("'%s'", v)
		}
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%F", v), "0"), ".")
	case bool:
		return fmt.Sprintf("%v", v)
	case time.Time:
		return fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05"))
	}

	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice || typ == reflect.Array {
		var s string
		vals := reflect.ValueOf(v)
		for i := 0; i < vals.Len(); i++ {
			if i != 0 {
				s += ", "
			}
			s += ToString(vals.Index(i).Interface(), true)
		}
		return s
	}

	return reflect.TypeOf(v).String()
}

// mapKeyType returns map key type with reflect.Type.
func mapKeyType(typ reflect.Type) reflect.Type {
	key := strings.Split(strings.Split(typ.String(), "[")[1], "]")[0]
	return typeStringToKind(key)
}

// mapValueType returns map value type with reflect.Type.
func mapValueType(typ reflect.Type) reflect.Type {
	val := strings.Split(typ.String(), "]")[1]
	return typeStringToKind(val)
}

func typeStringToKind(typ string) reflect.Type {
	switch typ {
	case "string":
		return reflect.TypeOf("")
	case "int":
		return reflect.TypeOf(0)
	case "int8":
		return reflect.TypeOf(int8(0))
	case "int16":
		return reflect.TypeOf(int16(0))
	case "int32":
		return reflect.TypeOf(int32(0))
	case "int64":
		return reflect.TypeOf(int64(0))
	case "uint":
		return reflect.TypeOf(uint(0))
	case "uint8":
		return reflect.TypeOf(uint8(0))
	case "uint16":
		return reflect.TypeOf(uint16(0))
	case "uint32":
		return reflect.TypeOf(uint32(0))
	case "uint64":
		return reflect.TypeOf(uint64(0))
	case "float32":
		return reflect.TypeOf(float32(0.0))
	case "float64":
		return reflect.TypeOf(float64(0.0))
	case "bool":
		return reflect.TypeOf(false)
	case "time.Time":
		return reflect.TypeOf(time.Time{})
	}
	return nil
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
