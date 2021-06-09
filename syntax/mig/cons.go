package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Cons is CONSTRAINT clause.
type Cons struct {
	Key string
}

// String returns function call as string.
func (c *Cons) String() string {
	return fmt.Sprintf("Cons(%s)", c.Key)
}

// Build creates the structure of CONSTRAINT clause that implements interfaces.ClauseSet.
func (c *Cons) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("CONSTRAINT")
	cs.WriteValue(c.Key)
	return cs, nil
}
