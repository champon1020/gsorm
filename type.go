package gsorm

import (
	"reflect"
	"time"

	"github.com/champon1020/gsorm/interfaces"
)

// conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type conn interface {
	Ping() error
	Query(query string, args ...interface{}) (irows, error)
	Exec(query string, args ...interface{}) (iresult, error)
}

type icolumnType interface {
	Name() string
	ScanType() reflect.Type
}

type irows interface {
	Next() bool
	Scan(args ...interface{}) error
	ColumnTypes() ([]icolumnType, error)
	Close() error
}

type iresult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// DB is the interface of database.
type DB interface {
	conn
	SetConnMaxLifetime(n time.Duration) error
	SetMaxIdleConns(n int) error
	SetMaxOpenConns(n int) error
	Close() error
	Begin() (Tx, error)
}

// Tx is the interface of database transaction.
type Tx interface {
	conn
	Commit() error
	Rollback() error
}

// Mock is mock database connection pool.
type Mock interface {
	conn
	compareWith(s interfaces.Stmt) (interface{}, error)
	Complete() error
	Expect(s interfaces.Stmt)
	ExpectWithReturn(s interfaces.Stmt, v interface{})
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
