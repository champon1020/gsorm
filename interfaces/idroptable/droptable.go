package idroptable

import "github.com/champon1020/mgorm/interfaces"

// Table is interface which is returned by mgorm.DropTable.
type Table interface {
	RawClause(rs string, v ...interface{}) RawClause
	interfaces.MigrateCallable
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	interfaces.MigrateCallable
}
