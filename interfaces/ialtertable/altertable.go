package ialtertable

import "github.com/champon1020/gsorm/interfaces"

// Table is interface which is returned by gsorm.AlterTable.
type Table interface {
	RawClause(rs string, v ...interface{}) RawClause
	Rename(string) Rename
	AddColumn(string, string) AddColumn
	DropColumn(string) DropColumn
	RenameColumn(string, string) RenameColumn
	AddCons(string) AddCons
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Rename(t string) Rename
	AddColumn(c string, t string) AddColumn
	NotNull() NotNull
	Default(v interface{}) Default
	DropColumn(c string) DropColumn
	RenameColumn(from string, to string) RenameColumn
	AddCons(k string) AddCons
	Unique(c ...string) Unique
	Primary(c ...string) Primary
	Foreign(c ...string) Foreign
	Ref(t string, c ...string) Ref
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
	Default(interface{}) Default
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
	Default(interface{}) Default
}

// Default is interface which is returned by (*AlterTableStmt).Default.
type Default interface {
	Table
	NotNull() NotNull
}

// AddCons is interface which is returned by (*AlterTableStmt).AddCons.
type AddCons interface {
	RawClause(rs string, v ...interface{}) RawClause
	Unique(...string) Unique
	Primary(...string) Primary
	Foreign(...string) Foreign
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
	RawClause(rs string, v ...interface{}) RawClause
	Ref(string, ...string) Ref
}

// Ref is interface which is returned by (*AlterTableStmt).Ref.
type Ref interface {
	Table
	interfaces.MigrateCallable
}
