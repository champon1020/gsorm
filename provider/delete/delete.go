package delete

import "github.com/champon1020/mgorm/provider"

// StmtMP is interface for returned value of mgorm.Delete.
type StmtMP interface {
	From(...string) FromMP
}

// FromMP is interface for returned value of (*Stmt).From.
type FromMP interface {
	Where(string, ...interface{}) WhereMP
	provider.ExecCallable
}

// WhereMP is interface for returned value of (*Stmt).Where.
type WhereMP interface {
	And(string, ...interface{}) AndMP
	Or(string, ...interface{}) OrMP
	provider.ExecCallable
}

// AndMP is interface for returned value of (*Stmt).And.
type AndMP interface {
	provider.ExecCallable
}

// OrMP is interface for returned value of (*Or).Or.
type OrMP interface {
	provider.ExecCallable
}
