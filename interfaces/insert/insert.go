package insert

import (
	"github.com/champon1020/mgorm/interfaces"
)

// Stmt is interface which is returned by mgorm.Insert.
type Stmt interface {
	Model(model interface{}) Model
	Select(selectStmt interfaces.QueryCallable) Select
	Values(vals ...interface{}) Values
}

// Model is interface which is returned by (*InsertStmt).Model.
type Model interface {
	interfaces.ExecCallable
}

// Select is interface which is returned by (*InsertStmt).Select.
type Select interface {
	interfaces.ExecCallable
}

// Values is interface which is returned by (*InsertStmt).Values.
type Values interface {
	Values(vals ...interface{}) Values
	interfaces.ExecCallable
}
