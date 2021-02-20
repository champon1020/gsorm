package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Delete is DELETE clause.
type Delete struct{}

// Name returns clause keyword.
func (d *Delete) Name() string {
	return "DELETE"
}

// String returns function call with string.
func (d *Delete) String() string {
	return fmt.Sprintf("%s()", d.Name())
}

// Build makes DELETE clause with syntax.StmtSet.
func (d *Delete) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Name())
	return ss, nil
}

// NewDelete create new delete object.
func NewDelete() *Delete {
	return &Delete{}
}
