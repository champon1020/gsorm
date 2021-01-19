package internal

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
			return r.String(), nil
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
