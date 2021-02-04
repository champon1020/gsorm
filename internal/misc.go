package internal

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Op values for error handling.
const (
	OpToString Op = "internal.toString"
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
func ToString(v interface{}) (string, error) {
	r := reflect.ValueOf(v)
	if r.IsValid() {
		switch r.Kind() {
		case reflect.String:
			s := fmt.Sprintf("%+q", r.String())
			return s, nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(r.Int(), 10), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(r.Uint(), 10), nil
		case reflect.Bool:
			return strconv.FormatBool(r.Bool()), nil
		}
	}
	return "", NewError(OpToString, KindType, errors.New("type is invalid"))
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
