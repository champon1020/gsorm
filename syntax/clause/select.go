package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Select is SELECT clause.
type Select struct {
	Columns []syntax.Column
}

// Keyword returns clause keyword.
func (s *Select) Keyword() string {
	return "SELECT"
}

// AddColumns appends column to Select.
func (s *Select) AddColumns(cols ...string) {
	for _, c := range cols {
		col := syntax.NewColumn(c)
		s.Columns = append(s.Columns, *col)
	}
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
	return fmt.Sprintf("%s(%s)", s.Keyword(), str)
}

// Build makes SELECT clause with syntax.StmtSet.
func (s *Select) Build() (domain.StmtSet, error) {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword(s.Keyword())
	for i, c := range s.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss, nil
}
