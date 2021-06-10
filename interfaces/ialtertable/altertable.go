package ialtertable

import "github.com/champon1020/gsorm/interfaces"

// Stmt is interface which is returned by gsorm.AlterTable.
type Stmt interface {
	RawClause(raw string, values ...interface{}) RawClause
	Rename(table string) Rename
	AddColumn(column, typ string) AddColumn
	DropColumn(column string) DropColumn
	RenameColumn(column, dest string) RenameColumn
	AddCons(key string) AddCons
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(raw string, values ...interface{}) RawClause
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
	Stmt
	interfaces.MigrateCallable
}

// AddColumn is interface which is returned by (*AlterTableStmt).AddColumn.
type AddColumn interface {
	NotNull() NotNull
	Default(value interface{}) Default
	Stmt
	interfaces.MigrateCallable
}

// DropColumn is interface which is returned by (*AlterTableStmt).DropColumn.
type DropColumn interface {
	Stmt
	interfaces.MigrateCallable
}

// RenameColumn is interface which is returned by (*AlterStmtStmt).RenameColumn.
type RenameColumn interface {
	Stmt
	interfaces.MigrateCallable
}

// NotNull is interface which is returned by (*AlterTableStmt).NotNull.
type NotNull interface {
	Default(value interface{}) Default
	Stmt
	interfaces.MigrateCallable
}

// Default is interface which is returned by (*AlterTableStmt).Default.
type Default interface {
	NotNull() NotNull
	Stmt
	interfaces.MigrateCallable
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
	Stmt
	interfaces.MigrateCallable
}

// Primary is interface which is returned by (*AlterTableStmt).Primary.
type Primary interface {
	Stmt
	interfaces.MigrateCallable
}

// Foreign is interface which is returned by (*AlterTableStmt).Foreign.
type Foreign interface {
	RawClause(raw string, values ...interface{}) RawClause
	Ref(table string, columns ...string) Ref
}

// Ref is interface which is returned by (*AlterTableStmt).Ref.
type Ref interface {
	Stmt
	interfaces.MigrateCallable
}
