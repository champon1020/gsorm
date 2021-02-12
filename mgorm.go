package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/champon1020/mgorm/syntax/cmd"
)

// New generate DB object.
func New(dn, dsn string) (*DB, error) {
	db, err := sql.Open(dn, dsn)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// NewMock generate MockDB object.
func NewMock() *MockDB {
	mock := new(MockDB)
	return mock
}

// Select statement api.
func Select(db Pool, cols ...string) SelectStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewSelect(cols)
	return stmt
}

// Insert statement api.
func Insert(db Pool, table string, cols ...string) InsertStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewInsert(table, cols)
	return stmt
}

// Update statement api.
func Update(db Pool, table string, cols ...string) UpdateStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewUpdate(table, cols)
	return stmt
}

// Delete statement api.
func Delete(db Pool) DeleteStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewDelete()
	return stmt
}

// Count statement api.
func Count(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Avg statement api.
func Avg(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Sum statement api.
func Sum(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Min statement api.
func Min(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Max statement api.
func Max(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// When statement api.
func When(e string, vals ...interface{}) WhenStmt {
	stmt := new(Stmt)
	stmt.call(clause.NewWhen(e, vals...))
	return stmt
}
