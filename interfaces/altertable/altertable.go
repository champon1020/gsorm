package alter

import "github.com/champon1020/mgorm/interfaces"

// TableMP is method interfaces which is returned by mgorm.AlterTable.
type TableMP interface {
	Rename(string) RenameMP
	AddColumn(string, string) AddColumnMP
	DropColumn(string) DropColumnMP
	RenameColumn(string, string) RenameColumnMP
	AddCons(string) AddConsMP
	DropPrimary(string) DropPrimaryMP
	DropForeign(string) DropForeignMP
	DropUnique(string) DropUniqueMP
}

// RenameMP is method interfaces which is returned by (*AlterTableStmt).Rename.
type RenameMP interface {
	TableMP
	interfaces.MigrateCallable
}

// AddColumnMP is method interfaces which is returned by (*AlterTableStmt).AddColumn.
type AddColumnMP interface {
	TableMP
	NotNull() NotNullMP
	AutoInc() AutoIncMP // Only MySQL
	Default(interface{}) DefaultMP
}

// DropColumnMP is method interfaces which is returned by (*AlterTableStmt).DropColumn.
type DropColumnMP interface {
	TableMP
	interfaces.MigrateCallable
}

// RenameColumnMP is method interfaces which is returned by (*AlterTableStmt).RenameColumn.
type RenameColumnMP interface {
	TableMP
	interfaces.MigrateCallable
}

// NotNullMP is method interfaces which is returned by (*AlterTableStmt).NotNull.
type NotNullMP interface {
	TableMP
	Default(interface{}) DefaultMP
	AutoInc() AutoIncMP
}

// DefaultMP is method interfaces which is returned by (*AlterTableStmt).Default.
type DefaultMP interface {
	TableMP
}

// AutoIncMP is method interfaces which is returned by (*AlterTableStmt).AutoInc.
type AutoIncMP interface {
	TableMP
}

// AddConsMP is method interfaces which is returned by (*AlterTableStmt).AddCons.
type AddConsMP interface {
	Unique(...string) UniqueMP
	Primary(...string) PrimaryMP
	Foreign(...string) ForeignMP
}

// DropUniqueMP is method interfaces which is returned by (*AlterTableStmt).DropUnique.
type DropUniqueMP interface {
	TableMP
	interfaces.MigrateCallable
}

// DropPrimaryMP is method interfaces which is returned by (*AlterTableStmt).DropPrimary.
type DropPrimaryMP interface {
	TableMP
	interfaces.MigrateCallable
}

// DropForeignMP is method interfaces which is returned by (*AlterTableStmt).DropForeign.
type DropForeignMP interface {
	TableMP
	interfaces.MigrateCallable
}

// UniqueMP is method interfaces which is returned by (*AlterTableStmt).Unique.
type UniqueMP interface {
	TableMP
	interfaces.MigrateCallable
}

// PrimaryMP is method interfaces which is returned by (*AlterTableStmt).Primary.
type PrimaryMP interface {
	TableMP
	interfaces.MigrateCallable
}

// ForeignMP is method interfaces which is returned by (*AlterTableStmt).Foreign.
type ForeignMP interface {
	Ref(string, string) RefMP
}

// RefMP is method interfaces which is returned by (*AlterTableStmt).Ref.
type RefMP interface {
	TableMP
	interfaces.MigrateCallable
}
