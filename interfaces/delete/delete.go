package delete

import "github.com/champon1020/mgorm/interfaces"

// StmtMP is interface for returned value of mgorm.Delete.
type StmtMP interface {
	From(...string) FromMP
}

// FromMP is interface for returned value of (*Stmt).From.
type FromMP interface {
	Where(string, ...interface{}) WhereMP
	interfaces.ExecCallable
}

// WhereMP is interface for returned value of (*Stmt).Where.
type WhereMP interface {
	And(string, ...interface{}) AndMP
	Or(string, ...interface{}) OrMP
	interfaces.ExecCallable
}

// AndMP is interface for returned value of (*Stmt).And.
type AndMP interface {
	interfaces.ExecCallable
}

// OrMP is interface for returned value of (*Or).Or.
type OrMP interface {
	interfaces.ExecCallable
}
