package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// On expression.
type On struct {
	Expr   string
	Values []interface{}
}

// Name returns string of clause.
func (o *On) Name() string {
	return "ON"
}

// String returns string of function call.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.Name(), s)
}

// Build make ON statement set.
func (o *On) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteClause(o.Name())
	return ss, nil
}

// NewOn create On instance.
func NewOn(expr string, vals ...interface{}) *On {
	return &On{Expr: expr, Values: vals}
}
