package syntax

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// toString convert the value of interface to string.
func toString(v interface{}) (string, error) {
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

	return "", newError(ErrInvalidType, "Type is not valid")
}

// PrintTestDiff prints the difference between expected and actual.
func PrintTestDiff(t *testing.T, diff string) {
	t.Errorf("Differs: (-got +want)\n%s", diff)
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func convertToSnakeCase(str string) (snake string) {
	str = strings.Split(str, "#")[0]
	snake = matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ToLower(snake)
	return
}
