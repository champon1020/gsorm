package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Having expression.
type Having struct {
	Expr   string
	Values []interface{}
}

// Name returns string of clause.
func (h *Having) Name() string {
	return "HAVING"
}

// String returns string of function call.
func (h *Having) String() string {
	s := fmt.Sprintf("%q", h.Expr)
	if len(h.Values) > 0 {
		s += ", "
		s += internal.SliceToString(h.Values)
	}
	return fmt.Sprintf("%s(%s)", h.Name(), s)
}

// Build makes HAVING statement set.
func (h *Having) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(h.Expr, h.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteClause(h.Name())
	return ss, nil
}

// NewHaving creates Having instance.
func NewHaving(expr string, vals ...interface{}) *Having {
	return &Having{Expr: expr, Values: vals}
}
