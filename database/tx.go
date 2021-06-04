package database

import (
	"database/sql"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/morikuni/failure"
)

// sqlTx is interface for sql.Tx.
type sqlTx interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// tx is an in-progress database transaction.
type tx struct {
	db   domain.DB
	conn sqlTx
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (t *tx) Ping() error {
	if t.db == nil {
		return failure.New(errFailedTxConnection, failure.Message("gsorm.tx.db is nil"))
	}
	return t.db.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (t *tx) Exec(query string, args ...interface{}) (domain.Result, error) {
	if t.conn == nil {
		return nil, failure.New(errFailedTxConnection, failure.Message("gsorm.tx.db is nil"))
	}
	r, err := t.conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &result{result: r}, nil
}

// Query executes a query that returns rows, typically a SELECT.
func (t *tx) Query(query string, args ...interface{}) (domain.Rows, error) {
	if t.conn == nil {
		return nil, failure.New(errFailedTxConnection, failure.Message("gsorm.tx.db is nil"))
	}
	r, err := t.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &rows{rows: r}, nil
}

// Commit commits the transaction.
func (t *tx) Commit() error {
	if t.conn == nil {
		return failure.New(errFailedTxConnection, failure.Message("gsorm.tx.db is nil"))
	}
	return t.conn.Commit()
}

// Rollback aborts the transaction.
func (t *tx) Rollback() error {
	if t.conn == nil {
		return failure.New(errFailedTxConnection, failure.Message("gsorm.tx.db is nil"))
	}
	return t.conn.Rollback()
}
