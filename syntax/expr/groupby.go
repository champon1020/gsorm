package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// GroupBy expression.
type GroupBy struct {
	Columns []syntax.Column
}

// Name returns string of clause.
func (g *GroupBy) Name() string {
	return "GROUP BY"
}

func (g *GroupBy) addColumn(col string) {
	g.Columns = append(g.Columns, *syntax.NewColumn(col))
}

// String returns string of function call.
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

// Build make GROUP BY statement set.
func (g *GroupBy) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(g.Name())
	for i, c := range g.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss, nil
}

// NewGroupBy creates GroupBy instance.
func NewGroupBy(cols []string) *GroupBy {
	g := new(GroupBy)
	for _, c := range cols {
		g.addColumn(c)
	}
	return g
}
