package drop

import "github.com/champon1020/mgorm/mp"

// DBMP is method provider which is returned by mgorm.DropDB.
type DBMP interface {
	mp.MigrationCallable
}

// IndexMP is method provider which is returned by mgorm.DropIndex.
type IndexMP interface {
	On(string) OnMP
	mp.MigrationCallable
}

// OnMP is method provider which is returned by (*DropIndex).On.
type OnMP interface {
	mp.MigrationCallable
}

// TableMP is method provider which is returned by mgorm.DropTable.
type TableMP interface {
	mp.MigrationCallable
}
