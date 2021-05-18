package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Set is SET clause.
type Set struct {
	Column string
	Value  interface{}
}

// Name returns clause keyword.
func (s *Set) Name() string {
	return "SET"
}

// String returns function call with string.
func (s *Set) String() string {
	return fmt.Sprintf("%s(%s, %v)", s.Name(), s.Column, s.Value)
}

// Build makes SET clause with syntax.StmtSet.
func (s *Set) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(s.Name())

	v := internal.ToString(s.Value, nil)
	ss.WriteValue(fmt.Sprintf("%s = %s", s.Column, v))
	return ss, nil
}
