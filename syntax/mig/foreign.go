package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Foreign is FOREIGN KEY clasue.
type Foreign struct {
	Columns []string
}

// String returns function call as string.
func (f *Foreign) String() string {
	return fmt.Sprintf("Foreign(%v)", f.Columns)
}

// Build creates the structure of FOREIGN KEY clause that implements interfaces.ClauseSet.
func (f *Foreign) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("FOREIGN KEY")
	cs.WriteValue("(")
	for i, c := range f.Columns {
		if i > 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(c)
	}
	cs.WriteValue(")")
	return cs, nil
}
