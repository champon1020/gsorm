package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Select is SELECT clause.
type Select struct {
	Columns []syntax.Column
}

// AddColumns appends the columns to Select.Columns.
func (s *Select) AddColumns(cols ...string) {
	for _, c := range cols {
		col := syntax.NewColumn(c)
		s.Columns = append(s.Columns, *col)
	}
}

// String returns function call as string.
func (s *Select) String() string {
	var str string
	for i, c := range s.Columns {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%q", c.Build())
	}
	return fmt.Sprintf("Select(%s)", str)
}

// Build creates the structure of SELECT clause that implements interfaces.ClauseSet.
func (s *Select) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("SELECT")
	for i, c := range s.Columns {
		if i != 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c.Build())
	}
	return cs, nil
}
