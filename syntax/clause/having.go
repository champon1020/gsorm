package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Having is HAVING clause.
type Having struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (h *Having) Name() string {
	return "HAVING"
}

// String returns function call with string.
func (h *Having) String() string {
	s := fmt.Sprintf("%q", h.Expr)
	if len(h.Values) > 0 {
		s += ", "
		s += internal.SliceToString(h.Values)
	}
	return fmt.Sprintf("%s(%s)", h.Name(), s)
}

// Build makes HAVING clause with syntax.StmtSet.
func (h *Having) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(h.Expr, h.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteKeyword(h.Name())
	return ss, nil
}
