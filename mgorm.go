package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/integration"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// New generate database object.
func New(db *sql.DB) *internal.Database {
	return &internal.Database{DB: db}
}

// NewMock generate MockDB object.
func NewMock() *integration.MockDB {
	mock := &integration.MockDB{Actual: []integration.QueryArgs{}}
	return mock
}

// Select statement api.
func Select(db internal.DB, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db internal.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db internal.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db internal.DB) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewDelete()
	return stmt
}
