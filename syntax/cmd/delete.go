package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Delete statement.
type Delete struct{}

// Query returns string of clause.
func (d *Delete) Query() string {
	return "DELETE"
}

// String returns string of function call.
func (d *Delete) String() string {
	return fmt.Sprintf("%s()", d.Query())
}

// Build make delete statement set.
func (d *Delete) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteClause(d.Query())
	return ss
}

// NewDelete create new delete object.
func NewDelete() *Delete {
	return &Delete{}
}
