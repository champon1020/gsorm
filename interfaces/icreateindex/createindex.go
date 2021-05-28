package icreateindex

import "github.com/champon1020/gsorm/interfaces"

// Index is interface which is returned by gsorm.CreateIndex.
type Index interface {
	RawClause(raw string, values ...interface{}) RawClause
	On(table string, columns ...string) On
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	On(table string, columns ...string) On
	interfaces.MigrateCallable
}

// On is interface which is returned by (*CreateIndexStmt).On.
type On interface {
	RawClause(raw string, values ...interface{}) RawClause
	interfaces.MigrateCallable
}
