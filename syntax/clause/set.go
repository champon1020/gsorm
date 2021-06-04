package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Set is SET clause.
type Set struct {
	Column string
	Value  interface{}
}

// String returns function call with string.
func (s *Set) String() string {
	return fmt.Sprintf("Set(%s, %v)", s.Column, s.Value)
}

// Build makes SET clause with syntax.StmtSet.
func (s *Set) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("SET")
	v := internal.ToString(s.Value, nil)
	ss.WriteValue(fmt.Sprintf("%s = %s", s.Column, v))
	return ss, nil
}
