package mgorm

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

// Exported values which is declared in stmt.go.
var (
	StmtProcessQuerySQL = (*Stmt).processQuerySQL
	StmtProcessExecSQL  = (*Stmt).processExecSQL
)
