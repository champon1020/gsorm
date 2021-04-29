package database

import (
	"database/sql"

	"github.com/champon1020/mgorm/internal"
)

// Conn is database connection like DB or Tx. This is also implemented by MockDB and MockTx.
type Conn interface {
	GetDriver() internal.SQLDriver
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
