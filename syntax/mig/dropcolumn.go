package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// DropColumn is DROP COLUMN clause.
type DropColumn struct {
	Column string
}

// String returns function call as string.
func (d *DropColumn) String() string {
	return fmt.Sprintf("DropColumn(%s)", d.Column)
}

// Build creates the structure of DROP COLUMN clause that implements interfaces.ClauseSet.
func (d *DropColumn) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("DROP COLUMN")
	cs.WriteValue(d.Column)
	return cs, nil
}
