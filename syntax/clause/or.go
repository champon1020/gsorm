package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Or is OR clause.
type Or struct {
	Expr   string
	Values []interface{}
}

// String returns function call with string.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("Or(%s)", s)
}

// Build makes OR clause with syntax.StmtSet.
func (o *Or) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword("OR")
	ss.Parens = true
	return ss, nil
}
