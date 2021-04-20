package delete

import "github.com/champon1020/mgorm/interfaces"

// Stmt is interface which is returned by mgorm.Delete.
type Stmt interface {
	From(...string) From
}

// From is interface which is returned by (*Stmt).From.
type From interface {
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*Stmt).Where.
type Where interface {
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*Stmt).And.
type And interface {
	interfaces.ExecCallable
}

// Or is interface which is returned by (*Or).Or.
type Or interface {
	interfaces.ExecCallable
}
