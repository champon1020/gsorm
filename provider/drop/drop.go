package drop

import "github.com/champon1020/mgorm/provider"

// DBMP is method provider which is returned by mgorm.DropDB.
type DBMP interface {
	provider.MigrationCallable
}

// IndexMP is method provider which is returned by mgorm.DropIndex.
type IndexMP interface {
	On(string) OnMP
	provider.MigrationCallable
}

// OnMP is method provider which is returned by (*DropIndex).On.
type OnMP interface {
	provider.MigrationCallable
}

// TableMP is method provider which is returned by mgorm.DropTable.
type TableMP interface {
	provider.MigrationCallable
}
