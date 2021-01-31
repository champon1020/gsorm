package mgorm

import (
	"database/sql"
	"errors"
	"time"
)

// DB is the db structure.
type DB struct {
	db *sql.DB
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (db *DB) Ping() error {
	if db.db == nil {
		return errors.New("DB is nil")
	}
	return db.db.Ping()
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (db *DB) SetConnMaxLifetime(n time.Duration) error {
	if db.db == nil {
		return errors.New("DB is nil")
	}
	db.db.SetConnMaxLifetime(n)
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (db *DB) SetMaxIdleConns(n int) error {
	if db.db == nil {
		return errors.New("DB is nil")
	}
	db.db.SetMaxIdleConns(n)
	return nil
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (db *DB) SetMaxOpenConns(n int) error {
	if db.db == nil {
		return errors.New("DB is nil")
	}
	db.db.SetMaxOpenConns(n)
	return nil
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (db *DB) Close() error {
	if db.db == nil {
		return errors.New("DB is nil")
	}
	return db.db.Close()
}
