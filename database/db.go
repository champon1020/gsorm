package database

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

// sqlDB is interface for sql.DB.
type sqlDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	SetConnMaxLifetime(n time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Close() error
	Begin() (*sql.Tx, error)
}

// DB is the interface of database.
type DB interface {
	Conn
	SetConnMaxLifetime(n time.Duration) error
	SetMaxIdleConns(n int) error
	SetMaxOpenConns(n int) error
	Close() error
	Begin() (Tx, error)
}

// db is a database handle representing a pool of zero or more underlying connections.
// It's safe for concurrent use by multiple goroutines.
type db struct {
	conn   sqlDB
	Driver internal.SQLDriver
}

// NewDB creates the DB instance.
func NewDB(dn, dsn string) (DB, error) {
	d, err := sql.Open(dn, dsn)
	if err != nil {
		return nil, err
	}
	if dn == "psql" {
		return &db{conn: d, Driver: internal.PSQL}, nil
	}
	return &db{conn: d, Driver: internal.MySQL}, nil
}

// GetDriver returns sql driver.
func (d *db) GetDriver() internal.SQLDriver {
	return d.Driver
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (d *db) Ping() error {
	if d.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	return d.conn.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	return d.conn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (d *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if d.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	return d.conn.Query(query, args...)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (d *db) SetConnMaxLifetime(n time.Duration) error {
	if d.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	d.conn.SetConnMaxLifetime(n)
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (d *db) SetMaxIdleConns(n int) error {
	if d.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	d.conn.SetMaxIdleConns(n)
	return nil
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (d *db) SetMaxOpenConns(n int) error {
	if d.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	d.conn.SetMaxOpenConns(n)
	return nil
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (d *db) Close() error {
	if d.conn == nil {
		return failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	return d.conn.Close()
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (d *db) Begin() (Tx, error) {
	if d.conn == nil {
		return nil, failure.New(errFailedDBConnection, failure.Message("mgorm.db.conn is nil"))
	}
	t, err := d.conn.Begin()
	if err != nil {
		return nil, failure.Wrap(err)
	}
	return &tx{db: d, conn: t}, nil
}
