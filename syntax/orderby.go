package syntax

// OrderBy expression.
type OrderBy struct {
	Column string
	Desc   bool
}

func (o *OrderBy) name() string {
	return "ORDER BY"
}

// Build make orderby statement set.
func (o *OrderBy) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(o.name())
	ss.WriteValue(o.Column)
	if o.Desc {
		ss.WriteValue("DESC")
	}
	return ss, nil
}

// NewOrderBy create new offset object.
func NewOrderBy(col string, desc bool) *OrderBy {
	return &OrderBy{Column: col, Desc: desc}
}
