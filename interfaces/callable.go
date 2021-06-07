package interfaces

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(model interface{}) error
	Stmt
}

// ExecCallable is embedded into clause interfaces which can call (*Stmt).Exec.
type ExecCallable interface {
	Exec() error
	Stmt
}

// MigrateCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrateCallable interface {
	Migrate() error
	SQL() string
}
