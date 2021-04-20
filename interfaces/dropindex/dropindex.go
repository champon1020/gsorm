package dropindex

import "github.com/champon1020/mgorm/interfaces"

// Index is interface which is returned by mgorm.DropIndex.
type Index interface {
	On(string) On
	interfaces.MigrateCallable
}

// On is interface which is returned by (*DropIndex).On.
type On interface {
	interfaces.MigrateCallable
}
