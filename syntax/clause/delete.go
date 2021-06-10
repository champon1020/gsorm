package clause

import (
	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Delete is DELETE clause.
type Delete struct{}

// String returns function call as string.
func (d *Delete) String() string {
	return "Delete()"
}

// Build creates the structure of DELETE clause that implements interfaces.ClauseSet.
func (d *Delete) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("DELETE")
	return cs, nil
}
