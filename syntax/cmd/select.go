package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Select statement.
type Select struct {
	Columns []syntax.Column
}

// Query returns string of clause.
func (s *Select) Query() string {
	return "SELECT"
}

func (s *Select) addColumn(col string) {
	c := syntax.NewColumn(col)
	s.Columns = append(s.Columns, *c)
}

// String returns string of function call.
func (s *Select) String() string {
	var str string
	for i, c := range s.Columns {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%q", c.Build())
	}
	return fmt.Sprintf("%s(%s)", s.Query(), str)
}

// Build make select statement set.
func (s *Select) Build() *syntax.StmtSet {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword(s.Query())
	for i, c := range s.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss
}

// NewSelect create new select object.
func NewSelect(cols []string) *Select {
	s := new(Select)
	for _, c := range cols {
		s.addColumn(c)
	}
	return s
}
