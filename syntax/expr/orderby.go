package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// OrderBy expression.
type OrderBy struct {
	Column string
	Desc   bool
}

// Name returns string of clause.
func (o *OrderBy) Name() string {
	return "ORDER BY"
}

// String returns string of function call.
func (o *OrderBy) String() string {
	return fmt.Sprintf("%s(%q, %v)", o.Name(), o.Column, o.Desc)
}

// Build make orderby statement set.
func (o *OrderBy) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(o.Name())
	ss.WriteValue(o.Column)
	if o.Desc {
		ss.WriteValue("DESC")
	}
	return ss, nil
}

// NewOrderBy create new offset object.
func NewOrderBy(col string, desc bool) *OrderBy {
	return &OrderBy{Column: col, Desc: desc}
}
