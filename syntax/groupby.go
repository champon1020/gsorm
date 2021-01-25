package syntax

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
