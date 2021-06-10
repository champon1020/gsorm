package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// CreateDB is CREATE DATABASE clause.
type CreateDB struct {
	DBName string
}

// String returns function call as string.
func (c *CreateDB) String() string {
	return fmt.Sprintf("CreateDB(%s)", c.DBName)
}

// Build creates the structure of CREATE DATABASE clause that implements interfaces.ClauseSet.
func (c *CreateDB) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("CREATE DATABASE")
	cs.WriteValue(c.DBName)
	return cs, nil
}
