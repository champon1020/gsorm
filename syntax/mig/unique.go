package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Unique is UNIQUE clause.
type Unique struct {
	Columns []string
}

// String returns function call as string.
func (u *Unique) String() string {
	return fmt.Sprintf("Unique(%v)", u.Columns)
}

// Build creates the structure of UNIQUE clause that implements interfaces.ClauseSet.
func (u *Unique) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("UNIQUE")
	cs.WriteValue("(")
	for i, c := range u.Columns {
		if i > 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c)
	}
	cs.WriteValue(")")
	return cs, nil
}
