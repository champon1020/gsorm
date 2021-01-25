package syntax

// Having expression.
type Having struct {
	Expr   string
	Values []interface{}
}

func (h *Having) name() string {
	return "HAVING"
}

// Build makes HAVING statement set.
func (h *Having) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(h.Expr, h.Values...)
	ss.WriteClause(h.name())
	return ss, err
}

// NewHaving creates Having instance.
func NewHaving(expr string, vals ...interface{}) *Having {
	return &Having{Expr: expr, Values: vals}
}
