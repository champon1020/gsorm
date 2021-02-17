package internal

// SQL string.
type SQL string

func (s *SQL) String() string {
	return string(*s)
}

func (s *SQL) Write(str string) {
	if len(*s) != 0 && str != ")" {
		*s += " "
	}
	*s += SQL(str)
}

func (s *SQL) Len() int {
	return len(s.String())
}
