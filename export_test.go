package gsorm

type (
	ExportedRawStmt     = rawStmt
	ExportedDB          = db
	ExportedTx          = tx
	ExportedMockDB      = mockDB
	ExportedIRows       = irows
	ExportedIResult     = iresult
	ExportedIColumnType = icolumnType
)

// Exported values which is declared in db.go.
func (d *db) ExportedSetConn(conn sqlDB) {
	d.conn = conn
}

func (t *tx) ExportedSetConn(conn sqlTx) {
	t.conn = conn
}

func (t *tx) ExportedSetDB(db *db) {
	t.db = db
}

func (s *stmt) ExportedGetErrors() []error {
	return s.errors
}

func (m *migStmt) ExportedGetErrors() []error {
	return m.errors
}
