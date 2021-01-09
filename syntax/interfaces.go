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

// DbIface is Db interface.
type DbIface interface {
	Query(string, ...interface{}) (RowsIface, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

// RowsIface is Rows interface.
type RowsIface interface {
	Close() error
	Columns() ([]string, error)
	Next() bool
	Scan(...interface{}) error
}
