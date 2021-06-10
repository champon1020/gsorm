package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// RenameColumn is RENAME COLUMN clause.
type RenameColumn struct {
	Column string
	Dest   string
}

// String returns function call as string.
func (r *RenameColumn) String() string {
	return fmt.Sprintf("RenameColumn(%s, %s)", r.Column, r.Dest)
}

// Build creates the structure of RENAME COLUMN clause that implements interfaces.ClauseSet.
func (r *RenameColumn) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("RENAME COLUMN")
	cs.WriteValue(r.Column)
	cs.WriteValue("TO")
	cs.WriteValue(r.Dest)
	return cs, nil
}
