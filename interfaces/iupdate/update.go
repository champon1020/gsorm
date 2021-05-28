package iupdate

import (
	"github.com/champon1020/gsorm/interfaces"
)

// Stmt is interface which is returned by gsorm.Update.
type Stmt interface {
	RawClause(raw string, values ...interface{}) RawClause
	Model(model interface{}, columns ...string) Model
	Set(column string, value interface{}) Set
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Set(column string, value interface{}) Set
	Where(expr string, values ...interface{}) Where
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
}

// Model is interface which is returned by (*UpdateStmt).Model.
type Model interface {
	Where(expr string, values ...interface{}) Where
	interfaces.ExecCallable
}

// Set is interface which is returned by (*UpdateStmt).Set.
type Set interface {
	RawClause(raw string, values ...interface{}) RawClause
	Set(column string, value interface{}) Set
	Where(epxr string, values ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*UpdateStmt).Where.
type Where interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*UpdateStmt).And.
type And interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}

// Or is interface which is returned by (*UpdateStmt).Or.
type Or interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}
