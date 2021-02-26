package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Stmt is interface for DeleteStmt, InsertStmt, SelectStmt, and so on.
type Stmt interface {
	String() string
	FuncString() string
	Called() []syntax.Clause
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

// Called gets called clauses.
func (s *stmt) Called() []syntax.Clause {
	return s.called
}
