package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Or is OR clause.
type Or struct {
	Expr   string
	Values []interface{}
}

// String returns function call as string.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("Or(%s)", s)
}

// Build creates the structure of OR clause that implements interfaces.ClauseSet.
func (o *Or) Build() (interfaces.ClauseSet, error) {
	s, err := syntax.BuildExpr(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	cs := &syntax.ClauseSet{Value: s}
	cs.WriteKeyword("OR")
	cs.Parens = true
	return cs, nil
}
