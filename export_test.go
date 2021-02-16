package mgorm

// Exported values which is declared in db.go.
func (db *DB) ExportedSetConn(conn sqlDB) {
	db.conn = conn
}

func (tx *Tx) ExportedSetConn(conn sqlTx) {
	tx.conn = conn
}

// Exported values which is declared in mockdb.go.
var (
	CompareStmts = compareStmts
)

// Exported values which is declared in stmt.go.
var (
	StmtProcessQuerySQL = (*Stmt).processQuerySQL
	StmtProcessCaseSQL  = (*Stmt).processCaseSQL
	StmtProcessExecSQL  = (*Stmt).processExecSQL
)

func (s *Stmt) ExportedGetErrors() []error {
	return s.errors
}
