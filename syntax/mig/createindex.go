package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// CreateIndex is CREATE INDEX clause.
type CreateIndex struct {
	IdxName string
}

// String returns function call as string.
func (c *CreateIndex) String() string {
	return fmt.Sprintf("CreateIndex(%s)", c.IdxName)
}

// Build creates the structure of CREATE INDEX clause that implements interfaces.ClauseSet.
func (c *CreateIndex) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("CREATE INDEX")
	cs.WriteValue(c.IdxName)
	return cs, nil
}
