package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Column is definition of table column.
type Column struct {
	Col  string
	Type string
}

// String returns function call as string.
func (c *Column) String() string {
	return fmt.Sprintf("Column(%s, %s)", c.Col, c.Type)
}

// Build creates the structure of column definition that implements interfaces.ClauseSet.
func (c *Column) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword(c.Col)
	cs.WriteValue(c.Type)
	return cs, nil
}
