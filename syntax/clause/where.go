package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Where is WHERE clause.
type Where struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (w *Where) Name() string {
	return "WHERE"
}

// String returns function call with string.
func (w *Where) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += ", "
		s += internal.ToString(w.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", w.Name(), s)
}

// Build makes WHERE clause with syntax.StmtSet.
func (w *Where) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(w.Expr, w.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(w.Name())
	return ss, nil
}
