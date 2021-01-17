package mgorm

import (
	"github.com/champon1020/mgorm/integration"
	"github.com/champon1020/mgorm/syntax"
)

// NewMock generate MockDb object.
func NewMock() *integration.MockDb {
	mock := &integration.MockDb{Actual: []integration.QueryArgs{}}
	return mock
}

// Select statement api.
func Select(db syntax.DB, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db syntax.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db syntax.DB, table string, cols ...string) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db syntax.DB) *Stmt {
	stmt := &Stmt{DB: db}
	stmt.Cmd = syntax.NewDelete()
	return stmt
}
