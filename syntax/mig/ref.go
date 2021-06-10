package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Ref is REFERENCES clause.
type Ref struct {
	Table   string
	Columns []string
}

// String returns function call as string.
func (r *Ref) String() string {
	return fmt.Sprintf("Ref(%s, %v)", r.Table, r.Columns)
}

// Build creates the structure of REFERENCES clause that implements interfaces.ClauseSet.
func (r *Ref) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("REFERENCES")
	cs.WriteValue(r.Table)
	cs.WriteValue("(")
	for i, c := range r.Columns {
		if i > 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c)
	}
	cs.WriteValue(")")
	return cs, nil
}
