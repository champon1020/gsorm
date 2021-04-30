package createtable

import "github.com/champon1020/mgorm/interfaces"

// Table is interface which is returned by mgorm.CreateTable.
type Table interface {
	Model(interface{}) Model
	Column(string, string) Column
}

// Model is interface which is returned by (*CreateTableStmt).Model.
type Model interface {
	interfaces.MigrateCallable
}

// Column is interface which is returned by (*CreateTableStmt).Column.
type Column interface {
	Column(string, string) Column
	NotNull() NotNull
	Default(interface{}) Default
	Cons(string) Cons
	interfaces.MigrateCallable
}

// NotNull is interface which is returned by (*CreateTableStmt).NotNull.
type NotNull interface {
	Column(string, string) Column
	Default(interface{}) Default
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Default is interface which is returned by (*CreateTableStmt).Default.
type Default interface {
	Column(string, string) Column
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Cons is interface which is returned by (*CreateTableStmt).Cons.
type Cons interface {
	Unique(...string) Unique
	Primary(...string) Primary
	Foreign(...string) Foreign
}

// Unique is interface which is returned by (*CreateTableStmt).Unique.
type Unique interface {
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Primary is interface which is returned by (*CreateTableStmt).Primary.
type Primary interface {
	Cons(string) Cons
	interfaces.MigrateCallable
}

// Foreign is interface which is returned by (*CreateTableStmt).Foreign.
type Foreign interface {
	Ref(string, ...string) Ref
}

// Ref is interface which is returned by (*CreateTableStmt).Ref.
type Ref interface {
	Cons(string) Cons
	interfaces.MigrateCallable
}
