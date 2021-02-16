package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

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

func (m *MigStmt) String() string {
	return ""
}

func (m *MigStmt) Migration() error {
	return nil
}

func (m *MigStmt) Column() ColumnMig {
	return m
}

func (m *MigStmt) NotNull() NotNullMig {
	return m
}

func (m *MigStmt) AutoInc() AutoIncMig {
	return m
}

func (m *MigStmt) Default(val interface{}) DefaultMig {
	return m
}

func (m *MigStmt) PK() PKMig {
	return m
}

func (m *MigStmt) FK(refCol string, refTable string) FKMig {
	return m
}
