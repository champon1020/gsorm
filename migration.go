package mgorm

import "github.com/champon1020/mgorm/syntax"

// MigStmt stores information about database migration query.
type MigStmt struct {
	pool   Pool
	cmd    syntax.MigCmd
	called []syntax.MigClause
	errors []error
}

// call appends called clause.
func (m *MigStmt) call(e syntax.MigClause) {
	m.called = append(m.called, e)
}

// throw appends occurred error.
func (m *MigStmt) throw(e error) {
	m.errors = append(m.errors, e)
}
