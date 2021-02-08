package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Exported values which is declared in mockdb.go.
var (
	MockDBCompareTo = (*MockDB).compareTo
)

func (m *MockDB) ExportedGetExpected() []*Stmt {
	return m.expected
}

func (m *MockDB) ExportedSetExpected(s []*Stmt) {
	m.expected = s
}

func (m *MockDB) ExportedGetWillReturn() map[int]interface{} {
	return m.willReturn
}

func (m *MockDB) ExportedSetWillReturn(willReturn map[int]interface{}) {
	m.willReturn = willReturn
}

// Exported values which is declared in stmt.go.
var (
	StmtProcessQuerySQL   = (*Stmt).processQuerySQL
	StmtProcessCaseSQL    = (*Stmt).processCaseSQL
	StmtProcessExecSQL    = (*Stmt).processExecSQL
	OpStmtProcessQuerySQL = opStmtProcessQuerySQL
	OpStmtProcessCaseSQL  = opStmtProcessCaseSQL
	OpStmtProcessExecSQL  = opStmtProcessExecSQL
)

func (s *Stmt) ExportedGetCmd() syntax.Cmd {
	return s.cmd
}

func (s *Stmt) ExportedSetCmd(cmd syntax.Cmd) {
	s.cmd = cmd
}

func (s *Stmt) ExportedGetCalled() []syntax.Expr {
	return s.called
}

func (s *Stmt) ExportedSetCalled(called []syntax.Expr) {
	s.called = called
}
