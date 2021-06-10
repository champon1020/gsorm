package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Where is WHERE clause.
type Where struct {
	Expr   string
	Values []interface{}
}

// String returns function call as string.
func (w *Where) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += ", "
		s += internal.ToString(w.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("Where(%s)", s)
}

// Build creates the structure of WHERE clause that implements interfaces.ClauseSet.
func (w *Where) Build() (interfaces.ClauseSet, error) {
	s, err := syntax.BuildExpr(w.Expr, w.Values...)
	if err != nil {
		return nil, err
	}
	cs := &syntax.ClauseSet{Value: s}
	cs.WriteKeyword("WHERE")
	return cs, nil
}
