package internal

// Op values for error handling.
const (
	opSQLDoQuery Op = "mgorm.SQL.doQuery"
	opSQLDoExec  Op = "mgorm.SQL.doExec"
	opSetField   Op = "mgorm.setField"
)

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
