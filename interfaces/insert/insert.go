package insert

import (
	"github.com/champon1020/mgorm/interfaces"
)

// StmtMP is interface for returned value of mgorm.Insert.
type StmtMP interface {
	Model(interface{}) ModelMP
	Select(interfaces.QueryCallable) SelectMP
	Values(...interface{}) ValuesMP
}

// ModelMP is interface for returned value of (*InsertStmt).Model.
type ModelMP interface {
	interfaces.ExecCallable
}

// SelectMP is interface for returned value of (*InsertStmt).Select.
type SelectMP interface {
	interfaces.ExecCallable
}

// ValuesMP is interface for returned value of (*InsertStmt).Values.
type ValuesMP interface {
	Values(...interface{}) ValuesMP
	interfaces.ExecCallable
}
