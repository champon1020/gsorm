package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// OrderBy is ORDER BY clause.
type OrderBy struct {
	Columns []string
}

// String returns function call with string.
func (o *OrderBy) String() string {
	return fmt.Sprintf("OrderBy(%q)", o.Columns)
}

// Build makes ORDER BY clause with sytnax.StmtSet.
func (o *OrderBy) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("ORDER BY")
	for i, c := range o.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	return ss, nil
}
