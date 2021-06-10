package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Set is SET clause.
type Set struct {
	Column string
	Value  interface{}
}

// String returns function call as string.
func (s *Set) String() string {
	return fmt.Sprintf("Set(%s, %v)", s.Column, s.Value)
}

// Build creates the structure of SET clause that implements interfaces.ClauseSet.
func (s *Set) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("SET")
	v := internal.ToString(s.Value, nil)
	cs.WriteValue(fmt.Sprintf("%s = %s", s.Column, v))
	return cs, nil
}
