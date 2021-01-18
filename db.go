package mgorm

import (
	"database/sql"
)

// DB is interface that is implemented by *sql.DB.
type DB interface {
	query(string, ...interface{}) (Rows, error)
	exec(string, ...interface{}) (sql.Result, error)
}

// Rows is interface that is implemented by *sql.Rows.
type Rows interface {
	Close() error
	Columns() ([]string, error)
	Next() bool
	Scan(...interface{}) error
}

// Database is the db structure.
type database struct {
	DB *sql.DB
}

// query executes the query that returns rows.
func (db *database) query(query string, args ...interface{}) (Rows, error) {
	return db.DB.Query(query, args...)
}

// exec executes without returning any rows.
func (db *database) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}
