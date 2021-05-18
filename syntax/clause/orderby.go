package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// OrderBy is ORDER BY clause.
type OrderBy struct {
	Columns []string
}

// Keyword returns clause keyword.
func (o *OrderBy) Keyword() string {
	return "ORDER BY"
}

// String returns function call with string.
func (o *OrderBy) String() string {
	return fmt.Sprintf("%s(%q)", o.Keyword(), o.Columns)
}

// Build makes ORDER BY clause with sytnax.StmtSet.
func (o *OrderBy) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Keyword())
	for i, c := range o.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	return ss, nil
}
