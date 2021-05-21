package icreateindex

import "github.com/champon1020/gsorm/interfaces"

// Index is interface which is returned by gsorm.CreateIndex.
type Index interface {
	RawClause(rs string, v ...interface{}) RawClause
	On(string, ...string) On
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	On(t string, c ...string) On
	interfaces.MigrateCallable
}

// On is interface which is returned by (*CreateIndexStmt).On.
type On interface {
	RawClause(rs string, v ...interface{}) RawClause
	interfaces.MigrateCallable
}
