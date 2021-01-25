package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Op values.
const (
	opSelect internal.Op = "mgorm.Select"
	opInsert internal.Op = "mgorm.Insert"
	opUpdate internal.Op = "mgorm.Update"
	opDelete internal.Op = "mgorm.Delete"
	opCount  internal.Op = "mgorm.Count"
	opAvg    internal.Op = "mgorm.Avg"
	opSum    internal.Op = "mgorm.Sum"
	opMin    internal.Op = "mgorm.Min"
	opMax    internal.Op = "mgorm.Max"
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
		called: []*opArgs{{op: opSelect, args: []interface{}{cols}}},
	}
	stmt.cmd = syntax.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db sqlDB, table string, cols ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opUpdate, args: []interface{}{table, cols}}},
	}
	stmt.cmd = syntax.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db sqlDB, table string, cols ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opUpdate, args: []interface{}{table, cols}}},
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

// Count statement api.
func Count(db sqlDB, col string, alias ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opCount, args: []interface{}{col}}},
	}
	s := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = syntax.NewSelect([]string{s})
	return stmt
}

// Avg statement api.
func Avg(db sqlDB, col string, alias ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opAvg, args: []interface{}{10}}},
	}
	s := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = syntax.NewSelect([]string{s})
	return stmt
}

// Sum statement api.
func Sum(db sqlDB, col string, alias ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opSum, args: []interface{}{10}}},
	}
	s := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = syntax.NewSelect([]string{s})
	return stmt
}

// Min statement api.
func Min(db sqlDB, col string, alias ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opMin, args: []interface{}{10}}},
	}
	s := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = syntax.NewSelect([]string{s})
	return stmt
}

// Max statement api.
func Max(db sqlDB, col string, alias ...string) *Stmt {
	stmt := &Stmt{
		db:     db,
		called: []*opArgs{{op: opMax, args: []interface{}{10}}},
	}
	s := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = syntax.NewSelect([]string{s})
	return stmt
}
