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
		s += internal.SliceToString(a.Values)
	}
	return fmt.Sprintf("%s(%s)", a.Name(), s)
}

// Build makes AND clause with syntax.StmtSet.
func (a *And) Build() (*syntax.StmtSet, error) {
	ss, err := syntax.BuildStmtSetForExpression(a.Expr, a.Values...)
	if err != nil {
		return nil, err
	}
	ss.WriteKeyword(a.Name())
	ss.Parens = true
	return ss, nil
}
