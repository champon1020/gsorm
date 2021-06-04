package gsorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/gsorm/interfaces/domain"
)

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
	CompareWith(domain.Stmt) (interface{}, error)
	Expect(s domain.Stmt)
	ExpectWithReturn(s domain.Stmt, v interface{})
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
