package update

import "github.com/champon1020/mgorm/provider"

// StmtMP is interface for returned value of mgorm.Update.
type StmtMP interface {
	Model(interface{}) ModelMP
	Set(...interface{}) SetMP
}

// ModelMP is interface for returned value of (*UpdateStmt).Model.
type ModelMP interface {
	Where(string, ...interface{}) WhereMP
	provider.ExecCallable
}

// SetMP is interface for returned value of (*UpdateStmt).Set.
type SetMP interface {
	Where(string, ...interface{}) WhereMP
	provider.ExecCallable
}

// WhereMP is interface for returned value of (*UpdateStmt).Where.
type WhereMP interface {
	And(string, ...interface{}) AndMP
	Or(string, ...interface{}) OrMP
	provider.ExecCallable
}

// AndMP is interface for returned value of (*UpdateStmt).And.
type AndMP interface {
	provider.ExecCallable
}

// OrMP is interface for returned value of (*UpdateStmt).Or.
type OrMP interface {
	provider.ExecCallable
}
