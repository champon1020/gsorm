package createtable

import "github.com/champon1020/mgorm/interfaces"

// TableMP is method provider which is returned by mgorm.CreateTable.
type TableMP interface {
	Model(interface{}) ModelMP
	Column(string, string) ColumnMP
}

// ModelMP is method provider which is returned by (*CreateTableStmt).Model.
type ModelMP interface {
	interfaces.MigrateCallable
}

// ColumnMP is method provider which is returned by (*CreateTableStmt).Column.
type ColumnMP interface {
	Column(string, string) ColumnMP
	NotNull() NotNullMP
	AutoInc() AutoIncMP // Only MySQL
	Default(interface{}) DefaultMP
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// NotNullMP is method provider which is returned by (*CreateTableStmt).NotNull.
type NotNullMP interface {
	Column(string, string) ColumnMP
	Default(interface{}) DefaultMP
	AutoInc() AutoIncMP
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// DefaultMP is method provider which is returned by (*CreateTableStmt).Default.
type DefaultMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// AutoIncMP is method provider which is returned by (*CreateTableStmt).AutoInc.
type AutoIncMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// ConsMP is method provider which is returned by (*CreateTableStmt).Cons.
type ConsMP interface {
	Unique(...string) UniqueMP
	Primary(...string) PrimaryMP
	Foreign(...string) ForeignMP
}

// UniqueMP is method provider which is returned by (*CreateTableStmt).Unique.
type UniqueMP interface {
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// PrimaryMP is method provider which is returned by (*CreateTableStmt).Primary.
type PrimaryMP interface {
	Cons(string) ConsMP
	interfaces.MigrateCallable
}

// ForeignMP is method provider which is returned by (*CreateTableStmt).Foreign.
type ForeignMP interface {
	Ref(string, string) RefMP
}

// RefMP is method provider which is returned by (*CreateTableStmt).Ref.
type RefMP interface {
	Cons(string) ConsMP
	interfaces.MigrateCallable
}
