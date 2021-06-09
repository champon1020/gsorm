package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Having is HAVING clause.
type Having struct {
	Expr   string
	Values []interface{}
}

// String returns function call as string.
func (h *Having) String() string {
	s := fmt.Sprintf("%q", h.Expr)
	if len(h.Values) > 0 {
		s += ", "
		s += internal.ToString(h.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("Having(%s)", s)
}

// Build creates the structure of HAVING clause that implements interfaces.ClauseSet.
func (h *Having) Build() (interfaces.ClauseSet, error) {
	s, err := syntax.BuildExpr(h.Expr, h.Values...)
	if err != nil {
		return nil, err
	}
	cs := &syntax.ClauseSet{Value: s}
	cs.WriteKeyword("HAVING")
	return cs, nil
}
