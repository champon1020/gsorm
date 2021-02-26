package provider

import "github.com/champon1020/mgorm/syntax"

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(interface{}) error
	String() string
	FuncString() string
	Called() []syntax.Clause
}

// ExecCallable is embedded into clause interfaces which can call (*Stmt).Exec.
type ExecCallable interface {
	Exec() error
	String() string
	FuncString() string
	Called() []syntax.Clause
}

// MigrationCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrationCallable interface {
	Migration() error
	String() string
}
