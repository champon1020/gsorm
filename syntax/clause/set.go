package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Set is SET clause.
type Set struct {
	Eqs []syntax.Eq
}

// Name returns clause keyword.
func (s *Set) Name() string {
	return "SET"
}

// AddEq appends equal expression to Set.
func (s *Set) AddEq(lhs string, rhs interface{}) {
	e := syntax.NewEq(lhs, rhs)
	s.Eqs = append(s.Eqs, *e)
}

// String returns function call with string.
func (s *Set) String() string {
	var str string
	for i, eq := range s.Eqs {
		if i != 0 {
			str += ", "
		}
		switch rhs := eq.RHS.(type) {
		case string:
			str += fmt.Sprintf("%q", rhs)
		default:
			str += fmt.Sprintf("%v", rhs)
		}
	}
	return fmt.Sprintf("%s(%s)", s.Name(), str)
}

// Build makes SET clause with syntax.StmtSet.
func (s *Set) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(s.Name())
	for i, eq := range s.Eqs {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(eq.Build())
	}
	return ss, nil
}
