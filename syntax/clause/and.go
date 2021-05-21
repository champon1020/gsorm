package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// And is AND clause.
type And struct {
	Expr   string
	Values []interface{}
}

// Keyword returns clause keyword.
func (a *And) Keyword() string {
	return "AND"
}

// String returns function call with string.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.ToString(a.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", a.Keyword(), s)
}

// Build makes AND clause with syntax.StmtSet.
func (a *And) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(a.Expr, a.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(a.Keyword())
	ss.Parens = true
	return ss, nil
}
