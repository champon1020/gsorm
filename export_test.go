package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Exported values which is declared in db.go.
func (db *DB) ExportedSetConn(conn sqlDB) {
	db.conn = conn
}

func (tx *Tx) ExportedSetConn(conn sqlTx) {
	tx.conn = conn
}

// Exported values which is declared in mockdb.go.
var (
	CompareStmts = compareStmts
)

// Exported values which is declared in stmt.go.
var (
	StmtProcessQuerySQL = (*Stmt).processQuerySQL
	StmtProcessCaseSQL  = (*Stmt).processCaseSQL
	StmtProcessExecSQL  = (*Stmt).processExecSQL
)

func (s *Stmt) ExportedGetCmd() syntax.Cmd {
	return s.cmd
}

func (s *Stmt) ExportedSetCmd(cmd syntax.Cmd) {
	s.cmd = cmd
}

func (s *Stmt) ExportedGetCalled() []syntax.Clause {
	return s.called
}

func (s *Stmt) ExportedSetCalled(called []syntax.Clause) {
	s.called = called
}

func (s *Stmt) ExportedGetErrors() []error {
	return s.errors
}
