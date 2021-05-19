package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Or is OR clause.
type Or struct {
	Expr   string
	Values []interface{}
}

// Keyword returns clause keyword.
func (o *Or) Keyword() string {
	return "OR"
}

// String returns function call with string.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", o.Keyword(), s)
}

// Build makes OR clause with syntax.StmtSet.
func (o *Or) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExpr(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(o.Keyword())
	ss.Parens = true
	return ss, nil
}
