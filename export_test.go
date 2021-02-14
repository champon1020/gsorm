package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// Exported values which is declared in mockdb.go.
func (m *MockDB) ExportedPushExpected(s *Stmt, v interface{}) {
	m.expected = append(m.expected, &expectedQuery{stmt: s, willReturn: v})
}

func (m *MockDB) ExportedPopExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	return m.expected[0]
}

func (m *MockTx) ExportedPushExpected(s *Stmt, v interface{}) {
	m.expected = append(m.expected, &expectedQuery{stmt: s, willReturn: v})
}

func (m *MockTx) ExportedPopExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	return m.expected[0]
}

type ExpectedQuery = expectedQuery

func (e *ExpectedQuery) ExportedGetStmt() *Stmt {
	return e.stmt
}

func (e *ExpectedQuery) ExportedSetStmt(s *Stmt) {
	e.stmt = s
}

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
