package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/internal"
)

// Conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type Conn interface {
	getDriver() internal.SQLDriver
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

// Mock is mock database conneciton pool.
type Mock interface {
	Conn
	Complete() error
	CompareWith(Stmt) (interface{}, error)
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
