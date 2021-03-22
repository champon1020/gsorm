package droptable

import "github.com/champon1020/mgorm/interfaces"

// TableMP is method interfaces which is returned by mgorm.DropTable.
type TableMP interface {
	interfaces.MigrateCallable
}
