package idropdb

import "github.com/champon1020/mgorm/interfaces"

// DB is interface which is returned by mgorm.DropDB.
type DB interface {
	interfaces.MigrateCallable
}
