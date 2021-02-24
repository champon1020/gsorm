package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Stmt is interface for DeleteStmt, InsertStmt, SelectStmt, and so on.
type Stmt interface {
	String() string
	funcString() string
	getCalled() []syntax.Clause
}

// stmt stores information about query.
type stmt struct {
	conn   Conn
	called []syntax.Clause
	model  interface{}
	errors []error
}

// call appends called clause.
func (s *stmt) call(e syntax.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *stmt) throw(err error) {
	s.errors = append(s.errors, err)
}

func (s *stmt) getCalled() []syntax.Clause {
	return s.called
}
