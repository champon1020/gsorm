package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/syntax"
)

// New generate database object.
func New(db *sql.DB) *database {
	return &database{DB: db}
}

// NewMock generate MockDB object.
func NewMock() *MockDB {
	mock := &MockDB{Actual: []queryArgs{}}
	return mock
}

// Select statement api.
func Select(db DB, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db DB) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewDelete()
	return stmt
}
