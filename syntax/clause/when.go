package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// When is WHEN clause.
type When struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (w *When) Name() string {
	return "WHEN"
}

// String returns function call with string.
func (w *When) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += fmt.Sprintf(", %s", internal.SliceToString(w.Values))
	}
	return fmt.Sprintf("%s(%s)", w.Name(), s)
}

// Build makes WHEN clause with syntax.StmtSet.
func (w *When) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(w.Expr, w.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteKeyword(w.Name())
	return ss, nil
}

// NewWhen creates When instance.
func NewWhen(expr string, vals ...interface{}) *When {
	return &When{Expr: expr, Values: vals}
}
