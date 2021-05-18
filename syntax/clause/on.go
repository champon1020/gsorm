package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// On is ON clause.
type On struct {
	Expr   string
	Values []interface{}
}

// Keyword returns clause keyword.
func (o *On) Keyword() string {
	return "ON"
}

// String returns function call with string.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", o.Keyword(), s)
}

// Build makes ON clause with syntax.StmtSet.
func (o *On) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExprWithoutQuotes(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(o.Keyword())
	return ss, nil
}
