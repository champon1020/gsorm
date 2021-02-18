package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/champon1020/mgorm/syntax/mig"
)

// New creates DB.
func New(dn, dsn string) (*DB, error) {
	db, err := sql.Open(dn, dsn)
	if err != nil {
		return nil, err
	}
	return &DB{conn: db, driver: dn}, nil
}

// NewMock creates MockDB.
func NewMock() *MockDB {
	mock := new(MockDB)
	return mock
}

// Select calls SELECT command.
func Select(db Pool, cols ...string) SelectStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewSelect(cols)
	return stmt
}

// Insert calls INSERT command.
func Insert(db Pool, table string, cols ...string) InsertStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewInsert(table, cols)
	return stmt
}

// Update calls UPDATE command.
func Update(db Pool, table string, cols ...string) UpdateStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewUpdate(table, cols)
	return stmt
}

// Delete calls DELETE command.
func Delete(db Pool) DeleteStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = cmd.NewDelete()
	return stmt
}

// Count calls COUNT function.
func Count(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Avg calls AVG function.
func Avg(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Sum calls SUM function.
func Sum(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Min calls MIN function.
func Min(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// Max calls MAX function.
func Max(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = cmd.NewSelect([]string{s})
	return stmt
}

// When calls CASE ... END clause.
func When(expr string, vals ...interface{}) WhenStmt {
	stmt := new(Stmt)
	stmt.call(&clause.When{Expr: expr, Values: vals})
	return stmt
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(pool Pool, dbName string) CreateDBMig {
	return &MigStmt{
		pool: pool,
		cmd:  &mig.CreateDB{DBName: dbName},
	}
}

// CreateTable calls CREATE TABLE command.
func CreateTable(pool Pool, table string) CreateTableMig {
	return &MigStmt{
		pool: pool,
		cmd:  &mig.CreateTable{Table: table},
	}
}

// AlterTable calls ALTER TABLE command.
func AlterTable(pool Pool, table string) AlterTableMig {
	return &MigStmt{
		pool: pool,
		cmd:  &mig.AlterTable{Table: table},
	}
}
