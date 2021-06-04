package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// GroupBy is GROUP BY clause.
type GroupBy struct {
	Columns []syntax.Column
}

// AddColumn appends column to GroupBy.
func (g *GroupBy) AddColumn(col string) {
	g.Columns = append(g.Columns, *syntax.NewColumn(col))
}

// String returns function call wtih string.
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

// Build makes GROUP BY clause with syntax.StmtSet.
func (g *GroupBy) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("GROUP BY")
	for i, c := range g.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss, nil
}
