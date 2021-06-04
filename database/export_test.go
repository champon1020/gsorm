package database

const (
	ErrInvalidMockExpectation = errInvalidMockExpectation
	ErrFailedDBConnection     = errFailedDBConnection
	ErrFailedTxConnection     = errFailedTxConnection
)

type ExportedDB = db
type ExportedTx = tx
type ExportedMockDB = mockDB

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
