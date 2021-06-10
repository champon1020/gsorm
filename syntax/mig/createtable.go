package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// CreateTable is CREATE TABLE clause.
type CreateTable struct {
	Table string
}

// String returns function call as string.
func (c *CreateTable) String() string {
	return fmt.Sprintf("CreateTable(%s)", c.Table)
}

// Build creates the structure of CREATE TABLE clause that implements interfaces.ClauseSet.
func (c *CreateTable) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("CREATE TABLE")
	cs.WriteValue(c.Table)
	return cs, nil
}
