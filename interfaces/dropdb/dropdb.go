package dropdb

import "github.com/champon1020/mgorm/interfaces"

// DBMP is method provider which is returned by mgorm.DropDB.
type DBMP interface {
	interfaces.MigrateCallable
}
