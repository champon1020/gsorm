package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Select is SELECT clause.
type Select struct {
	Columns []syntax.Column
}

// Name returns clause keyword.
func (s *Select) Name() string {
	return "SELECT"
}

// addColumn appends column to Select.
func (s *Select) addColumn(col string) {
	c := syntax.NewColumn(col)
	s.Columns = append(s.Columns, *c)
}

// String returns function call with string.
func (s *Select) String() string {
	var str string
	for i, c := range s.Columns {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%q", c.Build())
	}
	return fmt.Sprintf("%s(%s)", s.Name(), str)
}

// Build makes SELECT clause with syntax.StmtSet.
func (s *Select) Build() (*syntax.StmtSet, error) {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword(s.Name())
	for i, c := range s.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss, nil
}

// NewSelect create new select object.
func NewSelect(cols []string) *Select {
	s := new(Select)
	for _, c := range cols {
		s.addColumn(c)
	}
	return s
}
