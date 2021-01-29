package mgorm

import (
	"database/sql"
)

// DB is the db structure.
type DB struct {
	db *sql.DB
}

// query executes the query that returns rows.
func (db *DB) query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

// exec executes without returning any rows.
func (db *DB) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}
