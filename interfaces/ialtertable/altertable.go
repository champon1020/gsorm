package ialtertable

import "github.com/champon1020/gsorm/interfaces"

// Table is interface which is returned by gsorm.AlterTable.
type Table interface {
	RawClause(raw string, values ...interface{}) RawClause
	Rename(table string) Rename
	AddColumn(column, typ string) AddColumn
	DropColumn(column string) DropColumn
	RenameColumn(column, dest string) RenameColumn
	AddCons(key string) AddCons
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Rename(table string) Rename
	AddColumn(column string, typ string) AddColumn
	NotNull() NotNull
	Default(value interface{}) Default
	DropColumn(column string) DropColumn
	RenameColumn(column, dest string) RenameColumn
	AddCons(key string) AddCons
	Unique(columns ...string) Unique
	Primary(columns ...string) Primary
	Foreign(columns ...string) Foreign
	Ref(table string, columns ...string) Ref
}

// Rename is interface which is returned by (*AlterTableStmt).Rename.
type Rename interface {
	Table
	interfaces.MigrateCallable
}

// AddColumn is interface which is returned by (*AlterTableStmt).AddColumn.
type AddColumn interface {
	Table
	NotNull() NotNull
	Default(value interface{}) Default
}

// DropColumn is interface which is returned by (*AlterTableStmt).DropColumn.
type DropColumn interface {
	Table
	interfaces.MigrateCallable
}

// RenameColumn is interface which is returned by (*AlterTableStmt).RenameColumn.
type RenameColumn interface {
	Table
	interfaces.MigrateCallable
}

// NotNull is interface which is returned by (*AlterTableStmt).NotNull.
type NotNull interface {
	Table
	Default(value interface{}) Default
}

// Default is interface which is returned by (*AlterTableStmt).Default.
type Default interface {
	Table
	NotNull() NotNull
}

// AddCons is interface which is returned by (*AlterTableStmt).AddCons.
type AddCons interface {
	RawClause(raw string, values ...interface{}) RawClause
	Unique(columns ...string) Unique
	Primary(columns ...string) Primary
	Foreign(columns ...string) Foreign
}

// Unique is interface which is returned by (*AlterTableStmt).Unique.
type Unique interface {
	Table
	interfaces.MigrateCallable
}

// Primary is interface which is returned by (*AlterTableStmt).Primary.
type Primary interface {
	Table
	interfaces.MigrateCallable
}

// Foreign is interface which is returned by (*AlterTableStmt).Foreign.
type Foreign interface {
	RawClause(raw string, values ...interface{}) RawClause
	Ref(table string, columns ...string) Ref
}

// Ref is interface which is returned by (*AlterTableStmt).Ref.
type Ref interface {
	Table
	interfaces.MigrateCallable
}
