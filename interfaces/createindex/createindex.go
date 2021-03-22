package createindex

import "github.com/champon1020/mgorm/interfaces"

// IndexMP is method provider which is returned by mgorm.CreateIndex.
type IndexMP interface {
	On(string, ...string) OnMP
}

// OnMP is method provider which is returned by (*CreateIndexStmt).On.
type OnMP interface {
	interfaces.MigrateCallable
}
