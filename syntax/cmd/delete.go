package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Delete is DELETE clause.
type Delete struct{}

// Query returns clause keyword.
func (d *Delete) Query() string {
	return "DELETE"
}

// String returns function call with string.
func (d *Delete) String() string {
	return fmt.Sprintf("%s()", d.Query())
}

// Build makes DELETE clause with syntax.StmtSet.
func (d *Delete) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Query())
	return ss
}

// NewDelete create new delete object.
func NewDelete() *Delete {
	return &Delete{}
}
