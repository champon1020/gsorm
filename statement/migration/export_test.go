package migration

func (m *migStmt) ExportedGetErrors() []error {
	return m.errors
}
