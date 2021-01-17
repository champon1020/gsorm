package internal

import (
	"database/sql"
)

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

// Database is the db structure.
type Database struct {
	DB *sql.DB
}

// Query execute the query that returns rows.
func (db *Database) Query(query string, args ...interface{}) (Rows, error) {
	return db.DB.Query(query, args...)
}

// Exec execute without returning any rows.
func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}
