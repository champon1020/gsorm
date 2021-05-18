package icreatedb

import "github.com/champon1020/mgorm/interfaces"

// DB is interface which is returned by mgorm.CreateDB.
type DB interface {
	RawClause(rs string, v ...interface{}) RawClause
	interfaces.MigrateCallable
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	interfaces.MigrateCallable
}
