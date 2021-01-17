package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/integration"
	"github.com/champon1020/mgorm/syntax"
)

// NewMock generate MockDb object.
func NewMock() *integration.MockDb {
	mock := &integration.MockDb{Actual: []integration.QueryArgs{}}
	return mock
}

// Select statement api.
func Select(db *sql.DB, cols ...string) *Stmt {
	stmt := &Stmt{DB: wrapDB(db)}
	stmt.Cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db *sql.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: wrapDB(db)}
	stmt.Cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db *sql.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: wrapDB(db)}
	stmt.Cmd = syntax.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db *sql.DB) *Stmt {
	stmt := &Stmt{DB: wrapDB(db)}
	stmt.Cmd = syntax.NewDelete()
	return stmt
}
