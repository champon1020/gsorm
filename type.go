package gsorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/gsorm/interfaces/domain"
)

// conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type conn interface {
	Ping() error
	Query(query string, args ...interface{}) (irows, error)
	Exec(query string, args ...interface{}) (iresult, error)
}

type irows interface {
	Next() bool
	Scan(args ...interface{}) error
	ColumnTypes() ([]*sql.ColumnType, error)
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

// Mock is mock database conneciton pool.
type Mock interface {
	conn
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
