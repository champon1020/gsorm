package internal

import (
	"strconv"
	"time"
)

// Str is string type which can be converted to some types.
type Str string

// String converts s to string.
func (s Str) String() (string, error) {
	return string(s), nil
}

// Int converts s to int64.
func (s Str) Int() (int64, error) {
	return strconv.ParseInt(string(s), 10, 64)
}

// Uint converts s to uint64.
func (s Str) Uint() (uint64, error) {
	return strconv.ParseUint(string(s), 10, 64)
}

// Float converts s to float64.
func (s Str) Float() (float64, error) {
	return strconv.ParseFloat(string(s), 64)
}

// Bool converts s to bool.
func (s Str) Bool() (bool, error) {
	return strconv.ParseBool(string(s))
}

// Time converts s to time.Time.
func (s Str) Time(layout string) (time.Time, error) {
	return time.Parse(layout, string(s))
}

// Empty validates if s is empty or not.
func (s Str) Empty() bool {
	return s == ""
}
