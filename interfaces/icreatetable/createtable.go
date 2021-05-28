package icreatetable

import "github.com/champon1020/gsorm/interfaces"

// Table is interface which is returned by gsorm.CreateTable.
type Table interface {
	RawClause(raw string, values ...interface{}) RawClause
	Model(interface{}) Model
	Column(column, typ string) Column
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	Column(column, typ string) Column
	NotNull() NotNull
	Default(value interface{}) Default
	Cons(key string) Cons
	Unique(columns ...string) Unique
	Primary(columns ...string) Primary
	Foreign(columns ...string) Foreign
	Ref(table string, columns ...string) Ref
	interfaces.MigrateCallable
}

// Model is interface which is returned by (*CreateTableStmt).Model.
type Model interface {
	interfaces.MigrateCallable
}

// Column is interface which is returned by (*CreateTableStmt).Column.
type Column interface {
	RawClause(raw string, value ...interface{}) RawClause
	Column(column, typ string) Column
	NotNull() NotNull
	Default(value interface{}) Default
	Cons(key string) Cons
	interfaces.MigrateCallable
}

// NotNull is interface which is returned by (*CreateTableStmt).NotNull.
type NotNull interface {
	RawClause(raw string, value ...interface{}) RawClause
	Column(column, typ string) Column
	Default(value interface{}) Default
	Cons(key string) Cons
	interfaces.MigrateCallable
}

// Default is interface which is returned by (*CreateTableStmt).Default.
type Default interface {
	RawClause(raw string, value ...interface{}) RawClause
	Column(column, typ string) Column
	NotNull() NotNull
	Cons(key string) Cons
	interfaces.MigrateCallable
}

// Cons is interface which is returned by (*CreateTableStmt).Cons.
type Cons interface {
	RawClause(raw string, value ...interface{}) RawClause
	Unique(columns ...string) Unique
	Primary(columns ...string) Primary
	Foreign(columns ...string) Foreign
}

// Unique is interface which is returned by (*CreateTableStmt).Unique.
type Unique interface {
	RawClause(raw string, value ...interface{}) RawClause
	Cons(key string) Cons
	interfaces.MigrateCallable
}

// Primary is interface which is returned by (*CreateTableStmt).Primary.
type Primary interface {
	RawClause(raw string, value ...interface{}) RawClause
	Cons(key string) Cons
	interfaces.MigrateCallable
}

// Foreign is interface which is returned by (*CreateTableStmt).Foreign.
type Foreign interface {
	RawClause(raw string, value ...interface{}) RawClause
	Ref(table string, columns ...string) Ref
}

// Ref is interface which is returned by (*CreateTableStmt).Ref.
type Ref interface {
	RawClause(raw string, value ...interface{}) RawClause
	Cons(key string) Cons
	interfaces.MigrateCallable
}
