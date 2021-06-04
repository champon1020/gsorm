package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Where is WHERE clause.
type Where struct {
	Expr   string
	Values []interface{}
}

// String returns function call with string.
func (w *Where) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += ", "
		s += internal.ToString(w.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("Where(%s)", s)
}

// Build makes WHERE clause with syntax.StmtSet.
func (w *Where) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(w.Expr, w.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword("WHERE")
	return ss, nil
}
