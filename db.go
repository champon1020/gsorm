package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/errors"
)

// DB is a database handle representing a pool of zero or more underlying connections. It's safe for concurrent use by multiple goroutines.
type DB struct {
	conn *sql.DB
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (db *DB) Ping() error {
	if db.conn == nil {
		return errors.New("DB conn is nil", errors.InvalidValueError)
	}
	return db.conn.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (db *DB) SetConnMaxLifetime(n time.Duration) error {
	if db.conn == nil {
		return errors.New("DB conn is nil", errors.InvalidValueError)
	}
	db.conn.SetConnMaxLifetime(n)
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (db *DB) SetMaxIdleConns(n int) error {
	if db.conn == nil {
		return errors.New("DB conn is nil", errors.InvalidValueError)
	}
	db.conn.SetMaxIdleConns(n)
	return nil
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (db *DB) SetMaxOpenConns(n int) error {
	if db.conn == nil {
		return errors.New("DB conn is nil", errors.InvalidValueError)
	}
	db.conn.SetMaxOpenConns(n)
	return nil
}

// Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.
func (db *DB) Close() error {
	if db.conn == nil {
		return errors.New("DB conn is nil", errors.InvalidValueError)
	}
	return db.conn.Close()
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{db: db, conn: tx}, nil
}

// Tx is an in-progress database transaction.
type Tx struct {
	db   *DB
	conn *sql.Tx
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (tx *Tx) Ping() error {
	if tx.db == nil {
		return errors.New("Tx conn is nil", errors.InvalidValueError)
	}
	return tx.db.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.conn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (tx *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tx.conn.Query(query, args...)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	if tx.conn == nil {
		return errors.New("Tx conn is nil", errors.InvalidValueError)
	}
	return tx.conn.Commit()
}

// Rollback aborts the transaction.
func (tx *Tx) Rollback() error {
	if tx.conn == nil {
		return errors.New("Tx conn is nil", errors.InvalidValueError)
	}
	return tx.conn.Rollback()
}
