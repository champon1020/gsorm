package idropdb

import "github.com/champon1020/gsorm/interfaces"

// DB is interface which is returned by gsorm.DropDB.
type DB interface {
	RawClause(rs string, v ...interface{}) RawClause
	interfaces.MigrateCallable
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	interfaces.MigrateCallable
}
