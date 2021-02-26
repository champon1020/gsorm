package internal

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/champon1020/mgorm/errors"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ConvertToSnakeCase convert camel case string to snake case.
func ConvertToSnakeCase(str string) (snake string) {
	str = strings.Split(str, "#")[0]
	snake = matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ToLower(snake)
	return
}

// ToString convert the value of interface to string.
// If quotes is true, attache double quotes to string value.
func ToString(v interface{}, quotes bool) (string, error) {
	r := reflect.ValueOf(v)
	if r.IsValid() {
		switch r.Kind() {
		case reflect.String:
			if quotes {
				s := fmt.Sprintf("'%s'", r.String())
				return s, nil
			}
			return r.String(), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(r.Int(), 10), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(r.Uint(), 10), nil
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(r.Float(), 'f', -1, 64), nil
		case reflect.Bool:
			return strconv.FormatBool(r.Bool()), nil
		case reflect.Struct:
			t, ok := v.(time.Time)
			if ok {
				return fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05")), nil
			}
		default:
			return "", errors.New(fmt.Sprintf("Type %v is not supported", r.Kind()), errors.InvalidTypeError)
		}
	}
	return "", errors.New(fmt.Sprintf("Value %v is invalid", v), errors.InvalidValueError)
}

// SliceToString converts slice of interface{} to string separated with comma.
// If type of element is string, enclose with quotes.
func SliceToString(values []interface{}) string {
	var s string
	for i, v := range values {
		if i != 0 {
			s += ", "
		}
		switch v := v.(type) {
		case string:
			s += fmt.Sprintf("%q", v)
		default:
			s += fmt.Sprintf("%v", v)
		}
	}
	return s
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

// TimeFormat returns layout of date.
func TimeFormat(layout string) string {
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

// MapOfColumnsToFields returns map to localize between column and field.
func MapOfColumnsToFields(cols []string, modelTyp reflect.Type) map[int]int {
	indR2M := make(map[int]int)
	for i, c := range cols {
		for j := 0; j < modelTyp.NumField(); j++ {
			if c != ColumnNameFromTag(modelTyp.Field(j)) {
				continue
			}
			indR2M[i] = j
		}
	}
	return indR2M
}

// ColumnNameFromTag gets column name from struct field tag.
func ColumnNameFromTag(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return ConvertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}
