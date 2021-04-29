package database

import "github.com/champon1020/mgorm/internal"

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

func (d *db) ExportedSetDriver(driver internal.SQLDriver) {
	d.Driver = driver
}

func (t *tx) ExportedSetConn(conn sqlTx) {
	t.conn = conn
}

// Exported values which is declared in mockdb.go.
var (
	CompareStmts = compareStmts
)
