package database

import "github.com/champon1020/mgorm/domain"

const (
	ErrInvalidMockExpectation = errInvalidMockExpectation
	ErrFailedDBConnection     = errFailedDBConnection
	ErrFailedTxConnection     = errFailedTxConnection
)

type ExportedDB = db
type ExportedTx = tx
type ExportedMockDB = mockDB

// Exported values which is declared in mockdb.go.
var (
	CompareStmts = compareStmts
)

// Exported values which is declared in db.go.
func (d *db) ExportedSetConn(conn sqlDB) {
	d.conn = conn
}

func (d *db) ExportedSetDriver(driver domain.SQLDriver) {
	d.driver = driver
}

func (t *tx) ExportedSetConn(conn sqlTx) {
	t.conn = conn
}
