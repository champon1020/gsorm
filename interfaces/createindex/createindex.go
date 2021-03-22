package createindex

import "github.com/champon1020/mgorm/interfaces"

// Index is interface which is returned by mgorm.CreateIndex.
type Index interface {
	On(string, ...string) On
}

// On is interface which is returned by (*CreateIndexStmt).On.
type On interface {
	interfaces.MigrateCallable
}
