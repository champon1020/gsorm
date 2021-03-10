package create

import "github.com/champon1020/mgorm/provider"

// DBMP is method provider which is returned by mgorm.CreateDB.
type DBMP interface {
	provider.MigrateCallable
}

// IndexMP is method provider which is returned by mgorm.CreateIndex.
type IndexMP interface {
	On(string, ...string) OnMP
}

// OnMP is method provider which is returned by (*CreateIndexStmt).On.
type OnMP interface {
	provider.MigrateCallable
}

// TableMP is method provider which is returned by mgorm.CreateTable.
type TableMP interface {
	Model(interface{}) ModelMP
	Column(string, string) ColumnMP
}

// ModelMP is method provider which is returned by (*CreateTableStmt).Model.
type ModelMP interface {
	provider.MigrateCallable
}

// ColumnMP is method provider which is returned by (*CreateTableStmt).Column.
type ColumnMP interface {
	Column(string, string) ColumnMP
	NotNull() NotNullMP
	AutoInc() AutoIncMP // Only MySQL
	Default(interface{}) DefaultMP
	Cons(string) ConsMP
	provider.MigrateCallable
}

// NotNullMP is method provider which is returned by (*CreateTableStmt).NotNull.
type NotNullMP interface {
	Column(string, string) ColumnMP
	Default(interface{}) DefaultMP
	AutoInc() AutoIncMP
	Cons(string) ConsMP
	provider.MigrateCallable
}

// DefaultMP is method provider which is returned by (*CreateTableStmt).Default.
type DefaultMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	provider.MigrateCallable
}

// AutoIncMP is method provider which is returned by (*CreateTableStmt).AutoInc.
type AutoIncMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	provider.MigrateCallable
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
	provider.MigrateCallable
}

// PrimaryMP is method provider which is returned by (*CreateTableStmt).Primary.
type PrimaryMP interface {
	Cons(string) ConsMP
	provider.MigrateCallable
}

// ForeignMP is method provider which is returned by (*CreateTableStmt).Foreign.
type ForeignMP interface {
	Ref(string, string) RefMP
}

// RefMP is method provider which is returned by (*CreateTableStmt).Ref.
type RefMP interface {
	Cons(string) ConsMP
	provider.MigrateCallable
}
