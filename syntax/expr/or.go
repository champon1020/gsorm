package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Or clause.
type Or struct {
	Expr   string
	Values []interface{}
}

// Name returns string of clause.
func (o *Or) Name() string {
	return "OR"
}

// String returns string of function call.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.Name(), s)
}

// Build make OR statement set.
func (o *Or) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSet(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteClause(o.Name())
	ss.Parens = true
	return ss, nil
}

// NewOr create OR clause object.
func NewOr(expr string, vals ...interface{}) *Or {
	return &Or{Expr: expr, Values: vals}
}
