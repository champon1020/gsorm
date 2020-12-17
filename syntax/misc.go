package syntax

import (
	"reflect"
	"strconv"
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
