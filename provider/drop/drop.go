package drop

import "github.com/champon1020/mgorm/provider"

// DBMP is method provider which is returned by mgorm.DropDB.
type DBMP interface {
	provider.MigrateCallable
}

// IndexMP is method provider which is returned by mgorm.DropIndex.
type IndexMP interface {
	On(string) OnMP
	provider.MigrateCallable
}

// OnMP is method provider which is returned by (*DropIndex).On.
type OnMP interface {
	provider.MigrateCallable
}

// TableMP is method provider which is returned by mgorm.DropTable.
type TableMP interface {
	provider.MigrateCallable
}
