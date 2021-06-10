package gsorm

import (
	"database/sql"
	"time"

	"golang.org/x/xerrors"
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

type rows struct {
	rows *sql.Rows
}

func (r *rows) Next() bool {
	return r.rows.Next()
}

func (r *rows) Scan(args ...interface{}) error {
	return r.rows.Scan(args...)
}

func (r *rows) Close() error {
	return r.rows.Close()
}

func (r *rows) ColumnTypes() ([]icolumnType, error) {
	ct, err := r.rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	ict := make([]icolumnType, len(ct))
	for i, c := range ct {
		ict[i] = c
	}
	return ict, nil
}

type result struct {
	result sql.Result
}

func (r *result) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r *result) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

// db is a database handle representing a pool of zero or more underlying connections.
// It's safe for concurrent use by multiple goroutines.
type db struct {
	conn sqlDB
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (d *db) Ping() error {
	if d.conn == nil {
		return xerrors.New("gsorm.db.conn is nil")
	}
	return d.conn.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (d *db) Exec(query string, args ...interface{}) (iresult, error) {
	if d.conn == nil {
		return nil, xerrors.New("gsorm.db.conn is nil")
	}
	r, err := d.conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &result{result: r}, nil
}

// Query executes a query that returns rows, typically a SELECT.
func (d *db) Query(query string, args ...interface{}) (irows, error) {
	if d.conn == nil {
		return nil, xerrors.New("gsorm.db.conn is nil")
	}
	r, err := d.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &rows{rows: r}, nil
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (d *db) SetConnMaxLifetime(n time.Duration) error {
	if d.conn == nil {
		return xerrors.New("gsorm.db.conn is nil")
	}
	d.conn.SetConnMaxLifetime(n)
	return nil
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (d *db) SetMaxIdleConns(n int) error {
	if d.conn == nil {
		return xerrors.New("gsorm.db.conn is nil")
	}
	d.conn.SetMaxIdleConns(n)
	return nil
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (d *db) SetMaxOpenConns(n int) error {
	if d.conn == nil {
		return xerrors.New("gsorm.db.conn is nil")
	}
	d.conn.SetMaxOpenConns(n)
	return nil
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (d *db) Close() error {
	if d.conn == nil {
		return xerrors.New("gsorm.db.conn is nil")
	}
	return d.conn.Close()
}

// Begin starts a transaction. The default isolation level is dependent on the driver.
func (d *db) Begin() (Tx, error) {
	if d.conn == nil {
		return nil, xerrors.New("gsorm.db.conn is nil")
	}
	t, err := d.conn.Begin()
	if err != nil {
		return nil, err
	}
	return &tx{db: d, conn: t}, nil
}

// sqlTx is interface for sql.Tx.
type sqlTx interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// tx is an in-progress database transaction.
type tx struct {
	db   DB
	conn sqlTx
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (t *tx) Ping() error {
	if t.db == nil {
		return xerrors.New("gsorm.tx.conn is nil")
	}
	return t.db.Ping()
}

// Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.
func (t *tx) Exec(query string, args ...interface{}) (iresult, error) {
	if t.conn == nil {
		return nil, xerrors.New("gsorm.tx.conn is nil")
	}
	r, err := t.conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &result{result: r}, nil
}

// Query executes a query that returns rows, typically a SELECT.
func (t *tx) Query(query string, args ...interface{}) (irows, error) {
	if t.conn == nil {
		return nil, xerrors.New("gsorm.tx.conn is nil")
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
		return xerrors.New("gsorm.tx.conn is nil")
	}
	return t.conn.Commit()
}

// Rollback aborts the transaction.
func (t *tx) Rollback() error {
	if t.conn == nil {
		return xerrors.New("gsorm.tx.conn is nil")
	}
	return t.conn.Rollback()
}
