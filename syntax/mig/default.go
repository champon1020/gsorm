package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Default is DEFAULT clause.
type Default struct {
	Value interface{}
}

// String returns function call as string.
func (d *Default) String() string {
	return fmt.Sprintf("Default(%s)", d.Value)
}

// Build creates the structure of DEFAULT clause that implements interfaces.ClauseSet.
func (d *Default) Build() (interfaces.ClauseSet, error) {
	ss := &syntax.ClauseSet{}
	ss.WriteKeyword("DEFAULT")
	ss.WriteValue(internal.ToString(d.Value, nil))
	return ss, nil
}
