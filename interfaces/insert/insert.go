package insert

import (
	"github.com/champon1020/mgorm/interfaces"
)

// Stmt is interface which is returned by mgorm.Insert.
type Stmt interface {
	Model(interface{}) Model
	Select(interfaces.QueryCallable) Select
	Values(...interface{}) Values
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
	Values(...interface{}) Values
	interfaces.ExecCallable
}
