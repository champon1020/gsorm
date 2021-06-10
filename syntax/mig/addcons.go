package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// AddCons is ADD CONSTRAINT clause.
type AddCons struct {
	Key string
}

// String returns function call as string.
func (a *AddCons) String() string {
	return fmt.Sprintf("AddCons(%s)", a.Key)
}

// Build creates the structure of ADD CONSTRAINT clause that implements interfaces.ClauseSet.
func (a *AddCons) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("ADD CONSTRAINT")
	cs.WriteValue(a.Key)
	return cs, nil
}
