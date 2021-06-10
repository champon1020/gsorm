package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Primary is PRIMARY KEY clause.
type Primary struct {
	Columns []string
}

// String returns function call as string.
func (p *Primary) String() string {
	return fmt.Sprintf("Primary(%v)", p.Columns)
}

// Build creates the structure of PRIMARY KEY clause that implements interfaces.ClauseSet.
func (p *Primary) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("PRIMARY KEY")
	cs.WriteValue("(")
	for i, c := range p.Columns {
		if i > 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c)
	}
	cs.WriteValue(")")
	return cs, nil
}
