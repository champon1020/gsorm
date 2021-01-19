package mgorm

import (
	"database/sql"
)

// sqlDB is interface that is implemented by *sql.DB.
type sqlDB interface {
	query(string, ...interface{}) (sqlRows, error)
	exec(string, ...interface{}) (sql.Result, error)
}

// sqlRows is interface that is implemented by *sql.Rows.
type sqlRows interface {
	Close() error
	Columns() ([]string, error)
	Next() bool
	Scan(...interface{}) error
}

// DB is the db structure.
type DB struct {
	db *sql.DB
}

// query executes the query that returns rows.
func (db *DB) query(query string, args ...interface{}) (sqlRows, error) {
	return db.db.Query(query, args...)
}

// exec executes without returning any rows.
func (db *DB) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}
