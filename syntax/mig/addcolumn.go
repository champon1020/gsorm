package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// AddColumn is ADD clause.
type AddColumn struct {
	Column string
	Type   string
}

// String returns function call as string.
func (a *AddColumn) String() string {
	return fmt.Sprintf("AddColumn(%s, %s)", a.Column, a.Type)
}

// Build creates the structure of ADD COLUMN clause that implements interfaces.ClauseSet.
func (a *AddColumn) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("ADD COLUMN")
	cs.WriteValue(a.Column)
	cs.WriteValue(a.Type)
	return cs, nil
}
