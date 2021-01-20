package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Op values.
const (
	OpSelect internal.Op = "mgorm.Select"
	OpInsert internal.Op = "mgorm.Insert"
	OpUpdate internal.Op = "mgorm.Update"
	OpDelete internal.Op = "mgorm.OpDelete"
)

// New generate DB object.
func New(db *sql.DB) *DB {
	return &DB{db: db}
}

// NewMock generate MockDB object.
func NewMock() *MockDB {
	mock := new(MockDB)
	return mock
}

// Select statement api.
func Select(db sqlDB, cols ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: OpSelect, args: []interface{}{cols}}},
	}
	stmt.cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db sqlDB, table string, cols ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: OpUpdate, args: []interface{}{table, cols}}},
	}
	stmt.cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db sqlDB, table string, cols ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: OpUpdate, args: []interface{}{table, cols}}},
	}
	stmt.cmd = syntax.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db sqlDB) *Stmt {
	stmt := &Stmt{db: db}
	stmt.cmd = syntax.NewDelete()
	return stmt
}
