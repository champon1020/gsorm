package interfaces

import "github.com/champon1020/gsorm/interfaces/domain"

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(model interface{}) error
	domain.Stmt
}

// ExecCallable is embedded into clause interfaces which can call (*Stmt).Exec.
type ExecCallable interface {
	Exec() error
	domain.Stmt
}

// MigrateCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrateCallable interface {
	Migrate() error
	String() string
}
