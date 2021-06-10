package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// On is ON clause.
type On struct {
	Expr   string
	Values []interface{}
}

// String returns function call as string.
func (o *On) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.ToString(o.Values, &internal.ToStringOpt{DoubleQuotes: true})
	}
	return fmt.Sprintf("On(%s)", s)
}

// Build creates the structure of ON clause that implements interfaces.ClauseSet.
func (o *On) Build() (interfaces.ClauseSet, error) {
	s, err := syntax.BuildExprWithoutQuotes(o.Expr, o.Values...)
	if err != nil {
		return nil, err
	}
	cs := &syntax.ClauseSet{Value: s}
	cs.WriteKeyword("ON")
	return cs, nil
}
