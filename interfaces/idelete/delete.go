package idelete

import "github.com/champon1020/gsorm/interfaces"

// Stmt is interface which is returned by gsorm.Delete.
type Stmt interface {
	From(tables ...string) From
	RawClause(raw string, values ...interface{}) RawClause
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(raw string, values ...interface{}) RawClause
	From(tables ...string) From
	Where(expr string, values ...interface{}) Where
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}

// From is interface which is returned by (*Stmt).From.
type From interface {
	RawClause(raw string, values ...interface{}) RawClause
	Where(expr string, values ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*Stmt).Where.
type Where interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*Stmt).And.
type And interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	interfaces.ExecCallable
}

// Or is interface which is returned by (*Or).Or.
type Or interface {
	RawClause(raw string, values ...interface{}) RawClause
	Or(expr string, values ...interface{}) Or
	interfaces.ExecCallable
}
