package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// On is ON clause.
type On struct {
	Expr   string
	Values []interface{}
}

// String returns function call with string.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("On(%s)", s)
}

// Build makes ON clause with syntax.StmtSet.
func (o *On) Build() (domain.StmtSet, error) {
	s, err := syntax.BuildExprWithoutQuotes(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	ss := &syntax.StmtSet{Value: s}
	ss.WriteKeyword("ON")
	return ss, nil
}
