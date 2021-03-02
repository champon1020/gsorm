package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// And is AND clause.
type And struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (a *And) Name() string {
	return "AND"
}

// String returns function call with string.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.ToString(a.Values, true)
	}
	return fmt.Sprintf("%s(%s)", a.Name(), s)
}

// Build makes AND clause with syntax.StmtSet.
func (a *And) Build() (*syntax.StmtSet, error) {
	s, err := syntax.BuildForExpression(a.Expr, a.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(a.Name())
	ss.Parens = true
	return ss, nil
}
