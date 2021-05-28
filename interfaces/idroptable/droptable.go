package idroptable

import "github.com/champon1020/gsorm/interfaces"

// Table is interface which is returned by gsorm.DropTable.
type Table interface {
	RawClause(raw string, values ...interface{}) RawClause
	interfaces.MigrateCallable
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	interfaces.MigrateCallable
}
