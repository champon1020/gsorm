package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// OrderBy is ORDER BY clause.
type OrderBy struct {
	Column string
	Desc   bool
}

// Name returns clause keyword.
func (o *OrderBy) Name() string {
	return "ORDER BY"
}

// String returns function call with string.
func (o *OrderBy) String() string {
	return fmt.Sprintf("%s(%q, %v)", o.Name(), o.Column, o.Desc)
}

// Build makes ORDER BY clause with sytnax.StmtSet.
func (o *OrderBy) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Name())
	ss.WriteValue(o.Column)
	if o.Desc {
		ss.WriteValue("DESC")
	}
	return ss, nil
}
