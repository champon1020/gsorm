package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// OrderBy is ORDER BY clause.
type OrderBy struct {
	Columns []string
}

// Name returns clause keyword.
func (o *OrderBy) Name() string {
	return "ORDER BY"
}

// String returns function call with string.
func (o *OrderBy) String() string {
	return fmt.Sprintf("%s(%q)", o.Name(), o.Columns)
}

// Build makes ORDER BY clause with sytnax.StmtSet.
func (o *OrderBy) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Name())
	for i, c := range o.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	return ss, nil
}
