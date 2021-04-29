package migration

import (
	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/internal"
)

// Exported values which is declared in mig.go.
var (
	ExportedMySQLDB = &database.DB{Driver: internal.MySQL}
	ExportedPSQLDB  = &database.DB{Driver: internal.PSQL}
)

func (m *migStmt) ExportedGetErrors() []error {
	return m.errors
}
