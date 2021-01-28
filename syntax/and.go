package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

// And clause.
type And struct {
	Expr   string
	Values []interface{}
}

func (a *And) name() string {
	return "AND"
}

// String returns string of function call.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.SliceToString(a.Values)
	}
	return fmt.Sprintf("%s(%s)", a.name(), s)
}

// Build make AND statement set.
func (a *And) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(a.Expr, a.Values...)
	ss.WriteClause(a.name())
	ss.Parens = true
	return ss, err
}

// NewAnd create AND clause object.
func NewAnd(expr string, vals ...interface{}) *And {
	return &And{Expr: expr, Values: vals}
}
