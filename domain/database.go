package domain

import (
	"database/sql"
	"time"
)

// Conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type Conn interface {
	GetDriver() int
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
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

// Tx is the interface of database transaction.
type Tx interface {
	Conn
	Commit() error
	Rollback() error
}

// Mock is mock database conneciton pool.
type Mock interface {
	Conn
	Complete() error
	CompareWith(Stmt) (interface{}, error)
}

type MockDB interface {
	Mock
	SetConnMaxLifetime(n time.Duration) error
	SetMaxIdleConns(n int) error
	SetMaxOpenConns(n int) error
	Close() error
	Begin() (MockTx, error)
	ExpectBegin() MockTx
	Expect(stmt Stmt) MockDB
	Return(v interface{})
}

type MockTx interface {
	Mock
	Commit() error
	Rollback() error
	ExpectCommit()
	ExpectRollback()
	Expect(stmt Stmt) MockTx
	Return(v interface{})
}
