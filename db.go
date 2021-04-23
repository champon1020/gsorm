package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

// Conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type Conn interface {
	getDriver() internal.SQLDriver
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

// sqlDB is interface for sql.DB.
type sqlDB interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Ping() error
	SetConnMaxLifetime(n time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Close() error
	Begin() (*sql.Tx, error)
}

// sqlTx is interface for sql.Tx.
type sqlTx interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// DB is a database handle representing a pool of zero or more underlying connections. It's safe for concurrent use by multiple goroutines.
type DB struct {
	conn   sqlDB
	driver internal.SQLDriver
}

// getDriver returns sql driver.
func (db *DB) getDriver() internal.SQLDriver {
	return db.driver
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (db *DB) Ping() error {
	if db.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	return db.conn.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	return db.conn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	return db.conn.Query(query, args...)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (db *DB) SetConnMaxLifetime(n time.Duration) error {
	if db.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	db.conn.SetConnMaxLifetime(n)
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (db *DB) SetMaxIdleConns(n int) error {
	if db.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	db.conn.SetMaxIdleConns(n)
	return nil
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (db *DB) SetMaxOpenConns(n int) error {
	if db.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	db.conn.SetMaxOpenConns(n)
	return nil
}

// Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.
func (db *DB) Close() error {
	if db.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	return db.conn.Close()
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (db *DB) Begin() (*Tx, error) {
	if db.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.DB.conn is nil"))
	}
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, failure.Wrap(err)
	}
	return &Tx{db: db, conn: tx}, nil
}

// Tx is an in-progress database transaction.
type Tx struct {
	db   *DB
	conn sqlTx
}

func (tx *Tx) getDriver() internal.SQLDriver {
	return tx.db.driver
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (tx *Tx) Ping() error {
	if tx.db == nil {
		return failure.New(errFailedTxConnection, failure.Message("mgorm.Tx.db is nil"))
	}
	return tx.db.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if tx.conn == nil {
		return nil, failure.New(errFailedTxConnection, failure.Message("mgorm.Tx.db is nil"))
	}
	return tx.conn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (tx *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if tx.conn == nil {
		return nil, failure.New(errFailedTxConnection, failure.Message("mgorm.Tx.db is nil"))
	}
	return tx.conn.Query(query, args...)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	if tx.conn == nil {
		return failure.New(errFailedTxConnection, failure.Message("mgorm.Tx.db is nil"))
	}
	return tx.conn.Commit()
}

// Rollback aborts the transaction.
func (tx *Tx) Rollback() error {
	if tx.conn == nil {
		return failure.New(errFailedTxConnection, failure.Message("mgorm.Tx.db is nil"))
	}
	return tx.conn.Rollback()
}
