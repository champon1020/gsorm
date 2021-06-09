package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// NotNull is NOT NULL clause.
type NotNull struct{}

// String returns function call as string.
func (n *NotNull) String() string {
	return fmt.Sprintf("NotNull()")
}

// Build creates the structure of NOT NULL clause that implements interfaces.ClauseSet.
func (n *NotNull) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("NOT NULL")
	return cs, nil
}
