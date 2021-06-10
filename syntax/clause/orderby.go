package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// OrderBy is ORDER BY clause.
type OrderBy struct {
	Columns []string
}

// String returns function call as string.
func (o *OrderBy) String() string {
	return fmt.Sprintf("OrderBy(%q)", o.Columns)
}

// Build creates the structure of ORDER BY clause that implements interfaces.ClauseSet.
func (o *OrderBy) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("ORDER BY")
	for i, c := range o.Columns {
		if i > 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c)
	}
	return cs, nil
}
