package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

// On expression.
type On struct {
	Expr   string
	Values []interface{}
}

func (o *On) name() string {
	return "ON"
}

// String returns string of function call.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.name(), s)
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
