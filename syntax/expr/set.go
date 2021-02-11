package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/syntax"
)

// Set expression.
type Set struct {
	Eqs []syntax.Eq
}

// Name returns string of clause.
func (s *Set) Name() string {
	return "SET"
}

func (s *Set) addEq(lhs string, rhs interface{}) {
	e := syntax.NewEq(lhs, rhs)
	s.Eqs = append(s.Eqs, *e)
}

// String returns string of function call.
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

// Build make set statement set.
func (s *Set) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(s.Name())
	for i, eq := range s.Eqs {
		if i != 0 {
			ss.WriteValue(",")
		}
		e, err := eq.Build()
		if err != nil {
			return nil, err
		}
		ss.WriteValue(e)
	}
	return ss, nil
}

// NewSet create new set object.
func NewSet(lhs []string, rhs []interface{}) (*Set, error) {
	if len(lhs) != len(rhs) {
		return nil, errors.New("Length is different between lhs and rhs", errors.InvalidValueError)
	}
	s := new(Set)
	for i := 0; i < len(lhs); i++ {
		s.addEq(lhs[i], rhs[i])
	}
	return s, nil
}
