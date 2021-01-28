package syntax

import "fmt"

// GroupBy expression.
type GroupBy struct {
	Columns []Column
}

func (g *GroupBy) name() string {
	return "GROUP BY"
}

func (g *GroupBy) addColumn(col string) {
	g.Columns = append(g.Columns, *NewColumn(col))
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
	return fmt.Sprintf("%s(%s)", g.name(), s)
}

// Build make GROUP BY statement set.
func (g *GroupBy) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(g.name())
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
