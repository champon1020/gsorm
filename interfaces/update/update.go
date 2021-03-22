package update

import "github.com/champon1020/mgorm/interfaces"

// Stmt is interface which is returned by mgorm.Update.
type Stmt interface {
	Model(interface{}) Model
	Set(...interface{}) Set
}

// Model is interface which is returned by (*UpdateStmt).Model.
type Model interface {
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Set is interface which is returned by (*UpdateStmt).Set.
type Set interface {
	Where(string, ...interface{}) Where
	interfaces.ExecCallable
}

// Where is interface which is returned by (*UpdateStmt).Where.
type Where interface {
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	interfaces.ExecCallable
}

// And is interface which is returned by (*UpdateStmt).And.
type And interface {
	interfaces.ExecCallable
}

// Or is interface which is returned by (*UpdateStmt).Or.
type Or interface {
	interfaces.ExecCallable
}
