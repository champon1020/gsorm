package iupdate

import (
	"github.com/champon1020/mgorm/interfaces"
)

// Stmt is interface which is returned by mgorm.Update.
type Stmt interface {
	RawClause(rs string, v ...interface{}) RawClause
	Model(model interface{}, cols ...string) Model
	Set(col string, val interface{}) Set
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(rs string, v ...interface{}) RawClause
	Set(c string, v interface{}) Set
	Where(e string, v ...interface{}) Where
	And(e string, v ...interface{}) And
	Or(e string, v ...interface{}) Or
}

// Model is interface which is returned by (*UpdateStmt).Model.
type Model interface {
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Set is interface which is returned by (*UpdateStmt).Set.
type Set interface {
	RawClause(rs string, v ...interface{}) RawClause
	Set(col string, val interface{}) Set
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*UpdateStmt).Where.
type Where interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*UpdateStmt).And.
type And interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}

// Or is interface which is returned by (*UpdateStmt).Or.
type Or interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}
