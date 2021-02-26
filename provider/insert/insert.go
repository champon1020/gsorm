package insert

import (
	"github.com/champon1020/mgorm/provider"
)

// StmtMP is interface for returned value of mgorm.Insert.
type StmtMP interface {
	Model(interface{}) ModelMP
	Select(provider.QueryCallable) SelectMP
	Values(...interface{}) ValuesMP
}

// ModelMP is interface for returned value of (*InsertStmt).Model.
type ModelMP interface {
	provider.ExecCallable
}

// SelectMP is interface for returned value of (*InsertStmt).Select.
type SelectMP interface {
	provider.ExecCallable
}

// ValuesMP is interface for returned value of (*InsertStmt).Values.
type ValuesMP interface {
	Values(...interface{}) ValuesMP
	provider.ExecCallable
}
