package mgorm

import (
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Exported values which is declared in db.go.
func (db *DB) ExportedSetConn(conn sqlDB) {
	db.conn = conn
}

func (db *DB) ExportedSetDriver(driver internal.SQLDriver) {
	db.driver = driver
}

func (tx *Tx) ExportedSetConn(conn sqlTx) {
	tx.conn = conn
}

// Exported values which is declared in mig.go.
var (
	ExportedMySQLDB = &DB{driver: internal.MySQL}
	ExportedPSQLDB  = &DB{driver: internal.PSQL}
)

func (m *migStmt) ExportedGetErrors() []error {
	return m.errors
}

// Exported values which is declared in mockdb.go.
var (
	CompareStmts = compareStmts
)

// Exported values which is declared in stmt.go.
var (
	SelectStmtBuildSQL = (*SelectStmt).buildSQL
	InsertStmtBuildSQL = (*InsertStmt).buildSQL
	UpdateStmtBuildSQL = (*UpdateStmt).buildSQL
	DeleteStmtBuildSQL = (*DeleteStmt).buildSQL
)

func (s *stmt) ExportedGetErrors() []error {
	return s.errors
}

func (s *DeleteStmt) ExportedSetCalled(c ...syntax.Clause) {
	s.called = append(s.called, c...)
}

func (s *InsertStmt) ExportedSetCalled(c ...syntax.Clause) {
	s.called = append(s.called, c...)
}

func (s *SelectStmt) ExportedSetCalled(c ...syntax.Clause) {
	s.called = append(s.called, c...)
}

func (s *UpdateStmt) ExportedSetCalled(c ...syntax.Clause) {
	s.called = append(s.called, c...)
}
