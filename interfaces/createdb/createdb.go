package createdb

import "github.com/champon1020/mgorm/interfaces"

// DB is interface which is returned by mgorm.CreateDB.
type DB interface {
	interfaces.MigrateCallable
}
