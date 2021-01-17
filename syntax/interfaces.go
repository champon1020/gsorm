package syntax

import "database/sql"

// Cmd interface.
type Cmd interface {
	query() string
	Build() *StmtSet
}

// Expr interface.
type Expr interface {
	name() string
	Build() (*StmtSet, error)
}

// DB is interface that is implemented by *sql.DB.
type DB interface {
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

// Rows is interface that is implemented by *sql.Rows.
type Rows interface {
	Close() error
	Columns() ([]string, error)
	Next() bool
	Scan(...interface{}) error
}
