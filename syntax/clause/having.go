package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Having is HAVING clause.
type Having struct {
	Expr   string
	Values []interface{}
}

// Keyword returns clause keyword.
func (h *Having) Keyword() string {
	return "HAVING"
}

// String returns function call with string.
func (h *Having) String() string {
	s := fmt.Sprintf("%q", h.Expr)
	if len(h.Values) > 0 {
		s += ", "
		s += internal.ToString(h.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", h.Keyword(), s)
}

// Build makes HAVING clause with syntax.StmtSet.
func (h *Having) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(h.Expr, h.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(h.Keyword())
	return ss, nil
}
