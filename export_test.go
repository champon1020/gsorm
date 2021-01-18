package mgorm

import "database/sql"

// Exported values which is declared in stmt.go.
var (
	StmtProcessQuerySQL = (*Stmt).processQuerySQL
	StmtProcessExecSQL  = (*Stmt).processExecSQL
)

// Exported values which is declared in sql.go.
var (
	SQLString  = (*SQL).string
	SQLWrite   = (*SQL).write
	SQLDoQuery = (*SQL).doQuery
	SQLDoExec  = (*SQL).doExec
	SetToModel = setToModel
	ColumnName = columnName
	SetField   = setField
)

type TestMockDB struct {
	QueryFunc func(string, ...interface{}) (Rows, error)
	ExecFunc  func(string, ...interface{}) (sql.Result, error)
}

func (db *TestMockDB) query(query string, args ...interface{}) (Rows, error) {
	return db.QueryFunc(query, args...)
}
func (db *TestMockDB) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.ExecFunc(query, args...)
}

type TestMockRows struct {
	Max         int
	Count       int
	ColumnsFunc func() ([]string, error)
	ScanFunc    func(...interface{}) error
}

func (r *TestMockRows) Close() error { return nil }
func (r *TestMockRows) Columns() ([]string, error) {
	return r.ColumnsFunc()
}
func (r *TestMockRows) Next() bool {
	if r.Count >= r.Max {
		return false
	}
	r.Count++
	return true
}
func (r *TestMockRows) Scan(dest ...interface{}) error {
	return r.ScanFunc(dest...)
}
