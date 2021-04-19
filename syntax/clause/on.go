package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// On is ON clause.
type On struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (o *On) Name() string {
	return "ON"
}

// String returns function call with string.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, nil)
	}
	return fmt.Sprintf("%s(%s)", o.Name(), s)
}

// Build makes ON clause with syntax.StmtSet.
func (o *On) Build() (*syntax.StmtSet, error) {
	s, err := syntax.BuildForExpression(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword(o.Name())
	return ss, nil
}
