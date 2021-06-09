package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// On is ON clause which is used with CREATE INDEX.
type On struct {
	Table   string
	Columns []string
}

// String returns function call as string.
func (o *On) String() string {
	return fmt.Sprintf("On(%s, %v)", o.Table, o.Columns)
}

// Build creates the structure of ON clause that implements interfaces.ClauseSet.
func (o *On) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("ON")
	cs.WriteValue(o.Table)
	if len(o.Columns) > 0 {
		cs.WriteValue("(")
		for i, c := range o.Columns {
			if i > 0 {
				cs.WriteValue(",")
			}
			cs.WriteValue(c)
		}
		cs.WriteValue(")")
	}
	return cs, nil
}
