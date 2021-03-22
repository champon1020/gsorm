package dropindex

import "github.com/champon1020/mgorm/interfaces"

// IndexMP is method interfaces which is returned by mgorm.DropIndex.
type IndexMP interface {
	On(string) OnMP
	interfaces.MigrateCallable
}

// OnMP is method interfaces which is returned by (*DropIndex).On.
type OnMP interface {
	interfaces.MigrateCallable
}
