package idelete

import "github.com/champon1020/gsorm/interfaces"

// Stmt is interface which is returned by gsorm.Delete.
type Stmt interface {
	From(...string) From
	RawClause(rs string, v ...interface{}) RawClause
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	From(t ...string) From
	Where(e string, v ...interface{}) Where
	And(e string, v ...interface{}) And
	Or(e string, v ...interface{}) Or
	interfaces.ExecCallable
}

// From is interface which is returned by (*Stmt).From.
type From interface {
	RawClause(rs string, v ...interface{}) RawClause
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*Stmt).Where.
type Where interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*Stmt).And.
type And interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	interfaces.ExecCallable
}

// Or is interface which is returned by (*Or).Or.
type Or interface {
	RawClause(rs string, v ...interface{}) RawClause
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}
