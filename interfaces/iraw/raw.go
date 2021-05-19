package iraw

import "github.com/champon1020/mgorm/interfaces"

// Stmt is interface which is returned by mgorm.RawStmt.
type Stmt interface {
	interfaces.QueryCallable
	interfaces.ExecCallable
	interfaces.MigrateCallable
}
