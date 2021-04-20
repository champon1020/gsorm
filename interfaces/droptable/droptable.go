package droptable

import "github.com/champon1020/mgorm/interfaces"

// Table is interface which is returned by mgorm.DropTable.
type Table interface {
	interfaces.MigrateCallable
}
