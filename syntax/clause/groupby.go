package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// GroupBy is GROUP BY clause.
type GroupBy struct {
	Columns []syntax.Column
}

// AddColumn appends the column to GroupBy.Columns.
func (g *GroupBy) AddColumn(col string) {
	g.Columns = append(g.Columns, *syntax.NewColumn(col))
}

// String returns function call as string.
func (g *GroupBy) String() string {
	var s string
	for i, c := range g.Columns {
		if i != 0 {
			s += ", "
		}
		s += fmt.Sprintf("%q", c.Build())
	}
	return fmt.Sprintf("GroupBy(%s)", s)
}

// Build creates the structure of GROUP BY clause that implements interfaces.ClauseSet.
func (g *GroupBy) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("GROUP BY")
	for i, c := range g.Columns {
		if i != 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c.Build())
	}
	return cs, nil
}
