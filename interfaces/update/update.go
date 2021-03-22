package update

import "github.com/champon1020/mgorm/interfaces"

// StmtMP is interface for returned value of mgorm.Update.
type StmtMP interface {
	Model(interface{}) ModelMP
	Set(...interface{}) SetMP
}

// ModelMP is interface for returned value of (*UpdateStmt).Model.
type ModelMP interface {
	Where(string, ...interface{}) WhereMP
	interfaces.ExecCallable
}

// SetMP is interface for returned value of (*UpdateStmt).Set.
type SetMP interface {
	Where(string, ...interface{}) WhereMP
	interfaces.ExecCallable
}

// WhereMP is interface for returned value of (*UpdateStmt).Where.
type WhereMP interface {
	And(string, ...interface{}) AndMP
	Or(string, ...interface{}) OrMP
	interfaces.ExecCallable
}

// AndMP is interface for returned value of (*UpdateStmt).And.
type AndMP interface {
	interfaces.ExecCallable
}

// OrMP is interface for returned value of (*UpdateStmt).Or.
type OrMP interface {
	interfaces.ExecCallable
}
