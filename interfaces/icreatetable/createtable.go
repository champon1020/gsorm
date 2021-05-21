package icreatetable

import "github.com/champon1020/gsorm/interfaces"

// Table is interface which is returned by gsorm.CreateTable.
type Table interface {
	RawClause(rs string, v ...interface{}) RawClause
	Model(interface{}) Model
	Column(string, string) Column
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Column(c string, t string) Column
	NotNull() NotNull
	Default(v interface{}) Default
	Cons(k string) Cons
	Unique(c ...string) Unique
	Primary(c ...string) Primary
	Foreign(c ...string) Foreign
	Ref(t string, c ...string) Ref
	interfaces.MigrateCallable
}

// Model is interface which is returned by (*CreateTableStmt).Model.
type Model interface {
	interfaces.MigrateCallable
}

// Column is interface which is returned by (*CreateTableStmt).Column.
type Column interface {
	RawClause(rs string, v ...interface{}) RawClause
	Column(string, string) Column
	NotNull() NotNull
	Default(interface{}) Default
	Cons(string) Cons
	interfaces.MigrateCallable
}

// NotNull is interface which is returned by (*CreateTableStmt).NotNull.
type NotNull interface {
	RawClause(rs string, v ...interface{}) RawClause
	Column(string, string) Column
	Default(interface{}) Default
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Default is interface which is returned by (*CreateTableStmt).Default.
type Default interface {
	RawClause(rs string, v ...interface{}) RawClause
	Column(string, string) Column
	NotNull() NotNull
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Cons is interface which is returned by (*CreateTableStmt).Cons.
type Cons interface {
	RawClause(rs string, v ...interface{}) RawClause
	Unique(...string) Unique
	Primary(...string) Primary
	Foreign(...string) Foreign
}

// Unique is interface which is returned by (*CreateTableStmt).Unique.
type Unique interface {
	RawClause(rs string, v ...interface{}) RawClause
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Primary is interface which is returned by (*CreateTableStmt).Primary.
type Primary interface {
	RawClause(rs string, v ...interface{}) RawClause
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Foreign is interface which is returned by (*CreateTableStmt).Foreign.
type Foreign interface {
	RawClause(rs string, v ...interface{}) RawClause
	Ref(string, ...string) Ref
}

// Ref is interface which is returned by (*CreateTableStmt).Ref.
type Ref interface {
	RawClause(rs string, v ...interface{}) RawClause
	Cons(string) Cons
	interfaces.MigrateCallable
}
