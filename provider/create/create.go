package create

import "github.com/champon1020/mgorm/provider"

// DBMP is method provider which is returned by mgorm.CreateDB.
type DBMP interface {
	provider.MigrationCallable
}

// IndexMP is method provider which is returned by mgorm.CreateIndex.
type IndexMP interface {
	On(string, ...string) OnMP
}

// OnMP is method provider which is returned by (*CreateIndexStmt).On.
type OnMP interface {
	provider.MigrationCallable
}

// TableMP is method provider which is returned by mgorm.CreateTable.
type TableMP interface {
	Column(string, string) ColumnMP
}

// ColumnMP is method provider which is returned by (*CreateTableStmt).Column.
type ColumnMP interface {
	Column(string, string) ColumnMP
	NotNull() NotNullMP
	AutoInc() AutoIncMP // Only MySQL
	Default(interface{}) DefaultMP
	Cons(string) ConsMP
	provider.MigrationCallable
}

// NotNullMP is method provider which is returned by (*CreateTableStmt).NotNull.
type NotNullMP interface {
	Column(string, string) ColumnMP
	Default(interface{}) DefaultMP
	AutoInc() AutoIncMP
	Cons(string) ConsMP
	provider.MigrationCallable
}

// DefaultMP is method provider which is returned by (*CreateTableStmt).Default.
type DefaultMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	provider.MigrationCallable
}

// AutoIncMP is method provider which is returned by (*CreateTableStmt).AutoInc.
type AutoIncMP interface {
	Column(string, string) ColumnMP
	Cons(string) ConsMP
	provider.MigrationCallable
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
	provider.MigrationCallable
}

// PrimaryMP is method provider which is returned by (*CreateTableStmt).Primary.
type PrimaryMP interface {
	Cons(string) ConsMP
	provider.MigrationCallable
}

// ForeignMP is method provider which is returned by (*CreateTableStmt).Foreign.
type ForeignMP interface {
	Ref(string, string) RefMP
}

// RefMP is method provider which is returned by (*CreateTableStmt).Ref.
type RefMP interface {
	Cons(string) ConsMP
	provider.MigrationCallable
}
