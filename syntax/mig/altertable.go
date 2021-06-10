package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// AlterTable is ALTER TABLE clause.
type AlterTable struct {
	Table string
}

// String returns function call as string.
func (a *AlterTable) String() string {
	return fmt.Sprintf("AlterTable(%s)", a.Table)
}

// Build creates the structure of ALTER TABLE clause that implements interfaces.ClauseSet.
func (a *AlterTable) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("ALTER TABLE")
	cs.WriteValue(a.Table)
	return cs, nil
}
