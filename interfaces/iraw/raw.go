package iraw

import "github.com/champon1020/gsorm/interfaces"

// Stmt is interface which is returned by gsorm.RawStmt.
type Stmt interface {
	interfaces.QueryCallable
	interfaces.ExecCallable
	interfaces.MigrateCallable
}
