package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// DropTable is DROP TABLE clause.
type DropTable struct {
	Table string
}

// String returns function call as string.
func (d *DropTable) String() string {
	return fmt.Sprintf("DropTable(%s)", d.Table)
}

// Build creates the structure of DROP TABLE clause that implements interfaces.ClauseSet.
func (d *DropTable) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("DROP TABLE")
	cs.WriteValue(d.Table)
	return cs, nil
}
