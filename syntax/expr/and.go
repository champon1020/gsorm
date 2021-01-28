package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// And clause.
type And struct {
	Expr   string
	Values []interface{}
}

// Name returns string of clause.
func (a *And) Name() string {
	return "AND"
}

// String returns string of function call.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.SliceToString(a.Values)
	}
	return fmt.Sprintf("%s(%s)", a.Name(), s)
}

// Build make AND statement set.
func (a *And) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSet(a.Expr, a.Values...)
	ss.WriteClause(a.Name())
	ss.Parens = true
	return ss, err
}

// NewAnd create AND clause object.
func NewAnd(expr string, vals ...interface{}) *And {
	return &And{Expr: expr, Values: vals}
}
