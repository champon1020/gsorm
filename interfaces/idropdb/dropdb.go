package idropdb

import "github.com/champon1020/gsorm/interfaces"

// Stmt is interface which is returned by gsorm.DropDB.
type Stmt interface {
	RawClause(raw string, values ...interface{}) RawClause
	interfaces.MigrateCallable
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(raw string, values ...interface{}) RawClause
	interfaces.MigrateCallable
}
