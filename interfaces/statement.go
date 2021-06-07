package interfaces

import "github.com/champon1020/gsorm/interfaces/domain"

// Stmt is the interface for statements.
type Stmt interface {
	SQL() string
	String() string
	Clauses() []domain.Clause
	Cmd() domain.Clause
	CompareWith(s Stmt) error
}
