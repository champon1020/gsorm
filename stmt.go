package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Stmt stores information about query.
type Stmt struct {
	db     Pool
	called []syntax.Clause
	model  interface{}
	errors []error
}

// call appends called clause.
func (s *Stmt) call(e syntax.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *Stmt) throw(err error) {
	s.errors = append(s.errors, err)
}
