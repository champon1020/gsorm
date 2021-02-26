package alter

import "github.com/champon1020/mgorm/provider"

// TableMP is method provider which is returned by mgorm.AlterTable.
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

// RenameMP is method provider which is returned by (*AlterTableStmt).Rename.
type RenameMP interface {
	TableMP
	provider.MigrationCallable
}

// AddColumnMP is method provider which is returned by (*AlterTableStmt).AddColumn.
type AddColumnMP interface {
	TableMP
	NotNull() NotNullMP
	AutoInc() AutoIncMP // Only MySQL
	Default(interface{}) DefaultMP
}

// DropColumnMP is method provider which is returned by (*AlterTableStmt).DropColumn.
type DropColumnMP interface {
	TableMP
	provider.MigrationCallable
}

// RenameColumnMP is method provider which is returned by (*AlterTableStmt).RenameColumn.
type RenameColumnMP interface {
	TableMP
	provider.MigrationCallable
}

// NotNullMP is method provider which is returned by (*AlterTableStmt).NotNull.
type NotNullMP interface {
	TableMP
	Default(interface{}) DefaultMP
	AutoInc() AutoIncMP
}

// DefaultMP is method provider which is returned by (*AlterTableStmt).Default.
type DefaultMP interface {
	TableMP
}

// AutoIncMP is method provider which is returned by (*AlterTableStmt).AutoInc.
type AutoIncMP interface {
	TableMP
}

// AddConsMP is method provider which is returned by (*AlterTableStmt).AddCons.
type AddConsMP interface {
	Unique(...string) UniqueMP
	Primary(...string) PrimaryMP
	Foreign(...string) ForeignMP
}

// DropUniqueMP is method provider which is returned by (*AlterTableStmt).DropUnique.
type DropUniqueMP interface {
	TableMP
	provider.MigrationCallable
}

// DropPrimaryMP is method provider which is returned by (*AlterTableStmt).DropPrimary.
type DropPrimaryMP interface {
	TableMP
	provider.MigrationCallable
}

// DropForeignMP is method provider which is returned by (*AlterTableStmt).DropForeign.
type DropForeignMP interface {
	TableMP
	provider.MigrationCallable
}

// UniqueMP is method provider which is returned by (*AlterTableStmt).Unique.
type UniqueMP interface {
	TableMP
	provider.MigrationCallable
}

// PrimaryMP is method provider which is returned by (*AlterTableStmt).Primary.
type PrimaryMP interface {
	TableMP
	provider.MigrationCallable
}

// ForeignMP is method provider which is returned by (*AlterTableStmt).Foreign.
type ForeignMP interface {
	Ref(string, string) RefMP
}

// RefMP is method provider which is returned by (*AlterTableStmt).Ref.
type RefMP interface {
	TableMP
	provider.MigrationCallable
}
