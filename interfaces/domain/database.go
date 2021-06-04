package domain

import (
	"database/sql"
	"reflect"
	"time"
)

// SQLDriver is database driver.
type SQLDriver interface {
	LookupDefaultType(typ reflect.Type) string
}

// Conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type Conn interface {
	Ping() error
	Query(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) (Result, error)
}

type Rows interface {
	Next() bool
	Scan(args ...interface{}) error
	ColumnTypes() ([]*sql.ColumnType, error)
	Close() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
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
	Expect(s Stmt)
	ExpectWithReturn(s Stmt, v interface{})
}

// MockDB is interface of mock database.
type MockDB interface {
	Mock
	DB
	ExpectBegin() MockTx
}

// MockTx is interface of mock transaction.
type MockTx interface {
	Mock
	Tx
	ExpectCommit()
	ExpectRollback()
}
