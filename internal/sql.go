package internal

import "strings"

// SQL string.
type SQL string

func (s *SQL) String() string {
	return string(*s)
}

func (s *SQL) Write(str string) {
	if len(*s) != 0 &&
		str != ")" &&
		str != "," &&
		!strings.HasSuffix(s.String(), "(") {
		*s += " "
	}
	*s += SQL(str)
}

func (s *SQL) Len() int {
	return len(s.String())
}
