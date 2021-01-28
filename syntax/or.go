package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

// Or clause.
type Or struct {
	Expr   string
	Values []interface{}
}

func (o *Or) name() string {
	return "OR"
}

// String returns string of function call.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.name(), s)
}

// Build make OR statement set.
func (o *Or) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(o.Expr, o.Values...)
	ss.WriteClause(o.name())
	ss.Parens = true
	return ss, err
}

// NewOr create OR clause object.
func NewOr(expr string, vals ...interface{}) *Or {
	return &Or{Expr: expr, Values: vals}
}
