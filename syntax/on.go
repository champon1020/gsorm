package syntax

// On expression.
type On struct {
	Expr   string
	Values []interface{}
}

func (o *On) name() string {
	return "ON"
}

// Build make ON statement set.
func (o *On) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(o.Expr, o.Values...)
	ss.WriteClause(o.name())
	return ss, err
}

// NewOn create On instance.
func NewOn(expr string, vals ...interface{}) *On {
	return &On{Expr: expr, Values: vals}
}
