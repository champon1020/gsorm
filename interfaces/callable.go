package interfaces

import (
	"github.com/champon1020/mgorm/domain"
)

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(interface{}) error
	String() string
	FuncString() string
	Called() []domain.Clause
	Cmd() domain.Clause
	CompareWith(targetStmt domain.Stmt) error
}

// ExecCallable is embedded into clause interfaces which can call (*Stmt).Exec.
type ExecCallable interface {
	Exec() error
	String() string
	FuncString() string
	Called() []domain.Clause
	Cmd() domain.Clause
	CompareWith(targetStmt domain.Stmt) error
}

// MigrateCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrateCallable interface {
	Migrate() error
	String() string
}
