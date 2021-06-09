package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Rename is RENAME TO clause.
type Rename struct {
	Table string
}

// String returns function call as string.
func (r *Rename) String() string {
	return fmt.Sprintf("Rename(%s)", r.Table)
}

// Build creates the structure of RENAME TO clause that implements interfaces.ClauseSet.
func (r *Rename) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("RENAME TO")
	cs.WriteValue(r.Table)
	return cs, nil
}
