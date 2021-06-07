package iinsert

import (
	"github.com/champon1020/gsorm/interfaces"
)

// Stmt is interface which is returned by gsorm.Insert.
type Stmt interface {
	RawClause(raw string, values ...interface{}) RawClause
	Model(model interface{}) Model
	Select(stmt interfaces.Stmt) Select
	Values(values ...interface{}) Values
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Values(values ...interface{}) Values
	interfaces.ExecCallable
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
	RawClause(raw string, values ...interface{}) RawClause
	Values(values ...interface{}) Values
	interfaces.ExecCallable
}
