package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Or is OR clause.
type Or struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (o *Or) Name() string {
	return "OR"
}

// String returns function call with string.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.Name(), s)
}

// Build makes OR clause with syntax.StmtSet.
func (o *Or) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteKeyword(o.Name())
	ss.Parens = true
	return ss, nil
}

// NewOr create OR clause object.
func NewOr(expr string, vals ...interface{}) *Or {
	return &Or{Expr: expr, Values: vals}
}
