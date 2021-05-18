package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// GroupBy is GROUP BY clause.
type GroupBy struct {
	Columns []syntax.Column
}

// Name returns clause keyword.
func (g *GroupBy) Name() string {
	return "GROUP BY"
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
	return fmt.Sprintf("%s(%s)", g.Name(), s)
}

// Build makes GROUP BY clause with syntax.StmtSet.
func (g *GroupBy) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(g.Name())
	for i, c := range g.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss, nil
}
