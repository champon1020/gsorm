package mgorm

import (
	"database/sql"

	"github.com/champon1020/minigorm/syntax"
)

const (
	prodMode = 1 << 1
	testMode = 1 << 2
)

// Pool contains the connection pool.
type Pool struct {
	db   *sql.DB
	mode uint
}

// New creates the connection pool.
func New(drv string, dSrc string) *Pool {
	db, err := sql.Open(drv, dSrc)
	if err != nil {
		/* handle error */
	}

	return &Pool{db: db, mode: prodMode}
}

// Prod switch the mode property to production mode.
func (p *Pool) Prod() {
	p.mode = prodMode
}

// Test switch the mode property to test mode.
func (p *Pool) Test() {
	p.mode = testMode
}

// SetToStmt assigns Pool properties to Stmt.
func (p *Pool) setToStmt(stmt *syntax.Stmt, err ...error) {
	stmt.DB = p.db
	stmt.Mode = p.mode
	stmt.Errors = append(stmt.Errors, err...)
}

// Select statement api.
func (p *Pool) Select(cols []string, table []string) *syntax.Stmt {
	stmt := new(syntax.Stmt)
	stmt.Cmd = syntax.NewSelect(cols)
	stmt.From = syntax.NewFrom(table)
	p.setToStmt(stmt)
	return stmt
}

// Insert statement api.
func (p *Pool) Insert(table string, cols []string, vals []interface{}) *syntax.Stmt {
	stmt := new(syntax.Stmt)
	stmt.Cmd = syntax.NewInsert(table, cols)
	stmt.Values = syntax.NewValues(vals)
	p.setToStmt(stmt)
	return stmt
}

// Update statement api.
func (p *Pool) Update(table string, cols []string, vals []interface{}) *syntax.Stmt {
	stmt := new(syntax.Stmt)
	stmt.Cmd = syntax.NewUpdate(table)
	set, err := syntax.NewSet(cols, vals)
	stmt.Set = set
	p.setToStmt(stmt, err)
	return stmt
}

// Delete statement api.
func (p *Pool) Delete(table []string) *syntax.Stmt {
	stmt := new(syntax.Stmt)
	stmt.Cmd = syntax.NewDelete()
	stmt.From = syntax.NewFrom(table)
	p.setToStmt(stmt)
	return stmt
}
