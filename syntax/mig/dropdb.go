package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// DropDB is DROP DATABASE clause.
type DropDB struct {
	DBName string
}

// String returns function call as string.
func (d *DropDB) String() string {
	return fmt.Sprintf("DropDB(%s)", d.DBName)
}

// Build creates the structure of DROP DATABASE clause that implements interfaces.ClauseSet.
func (d *DropDB) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("DROP DATABASE")
	cs.WriteValue(d.DBName)
	return cs, nil
}
