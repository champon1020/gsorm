package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// And is AND clause.
type And struct {
	Expr   string
	Values []interface{}
}

// String returns function call as string.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.ToString(a.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("And(%s)", s)
}

// Build creates the structure of AND clause that implements interfaces.ClauseSet.
func (a *And) Build() (interfaces.ClauseSet, error) {
	s, err := syntax.BuildExpr(a.Expr, a.Values...)
	if err != nil {
		return nil, err
	}
	cs := &syntax.ClauseSet{Value: s}
	cs.WriteKeyword("AND")
	cs.Parens = true
	return cs, nil
}
