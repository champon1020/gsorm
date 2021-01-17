package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/syntax"
)

type database struct {
	DB *sql.DB
}

func (db *database) Query(query string, args ...interface{}) (syntax.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db *database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func wrapDB(db *sql.DB) *database {
	return &database{DB: db}
}
