package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Pool is database connection pool like DB or Tx. This is also implemented by MockDB and MockTx.
type Pool interface {
	getDriver() internal.SQLDriver
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

// Mock is mock database conneciton pool.
type Mock interface {
	Pool
	Complete() error
	CompareWith(*Stmt) (interface{}, error)
}

// sqlDB is interface for sql.DB.
type sqlDB interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Ping() error
	SetConnMaxLifetime(n time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Close() error
	Begin() (*sql.Tx, error)
}

// sqlTx is interface for sql.Tx.
type sqlTx interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(interface{}) error
	ExpectQuery(interface{}) *Stmt
	String() string
}

// ExecCallable is embedded into clause interfaces which can call (*Stmt).Exec.
type ExecCallable interface {
	Exec() error
	ExpectExec() *Stmt
	String() string
}

// SelectStmt is returned after Select is called.
type SelectStmt interface {
	From(...string) FromStmt
}

// InsertStmt is returned after Insert is called.
type InsertStmt interface {
	Values(...interface{}) ValuesStmt
	Model(interface{}) ModelStmt
}

// UpdateStmt is returned after Update is called.
type UpdateStmt interface {
	Set(...interface{}) SetStmt
	Model(interface{}) ModelStmt
}

// DeleteStmt is returned after Delete is called.
type DeleteStmt interface {
	From(...string) FromStmt
}

// ModelStmt is returned after (*Stmt).Model is called.
type ModelStmt interface {
	Where(string, ...interface{}) WhereStmt
	ExecCallable
}

// FromStmt is returned after (*Stmt).From is called.
type FromStmt interface {
	Join(string) JoinStmt
	LeftJoin(string) JoinStmt
	RightJoin(string) JoinStmt
	FullJoin(string) JoinStmt
	Where(string, ...interface{}) WhereStmt
	GroupBy(...string) GroupByStmt
	OrderBy(...string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// ValuesStmt is returned after (*Stmt).Values is called.
type ValuesStmt interface {
	ExecCallable
}

// SetStmt is returned after (*Stmt).Set is called.
type SetStmt interface {
	Where(string, ...interface{}) WhereStmt
	ExecCallable
}

// JoinStmt is returned after (*Stmt).Join, (*Stmt).LeftJoin, (*Stmt).RightJoin or (*Stmt).FullJoin is called.
type JoinStmt interface {
	On(string, ...interface{}) OnStmt
}

// OnStmt is returned after (*Stmt).On is called.
type OnStmt interface {
	Where(string, ...interface{}) WhereStmt
	GroupBy(...string) GroupByStmt
	OrderBy(...string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// WhereStmt is returned after (*Stmt).Where is called.
type WhereStmt interface {
	And(string, ...interface{}) AndStmt
	Or(string, ...interface{}) OrStmt
	GroupBy(...string) GroupByStmt
	OrderBy(...string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
	ExecCallable
}

// AndStmt is returned after (*Stmt).And is called.
type AndStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(...string) OrderByStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
	ExecCallable
}

// OrStmt is returned after (*Stmt).Or is called.
type OrStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(...string) OrderByStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
	ExecCallable
}

// GroupByStmt is returned after (*Stmt).GroupBy is called.
type GroupByStmt interface {
	Having(string, ...interface{}) HavingStmt
	OrderBy(...string) OrderByStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// HavingStmt is returned after (*Stmt).Having is called.
type HavingStmt interface {
	OrderBy(...string) OrderByStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// OrderByStmt is returned after (*Stmt).OrderBy is called.
type OrderByStmt interface {
	Limit(int) LimitStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// LimitStmt is returned after (*Stmt).Limit is called.
type LimitStmt interface {
	Offset(int) OffsetStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// OffsetStmt is returned after (*Stmt).Offset is called.
type OffsetStmt interface {
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// UnionStmt is returned after (*Stmt).Union is called.
type UnionStmt interface {
	OrderBy(...string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Stmt) UnionStmt
	UnionAll(syntax.Stmt) UnionStmt
	QueryCallable
}

// MigrationCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrationCallable interface {
	Migration() error
	String() string
}

// CreateDBMig is returned after CreateDB is called.
type CreateDBMig interface {
	MigrationCallable
}

// DropDBMig is returned after DropDB is called.
type DropDBMig interface {
	MigrationCallable
}

// CreateTableMig is returned after CreateTable is called.
type CreateTableMig interface {
	Column(string, string) ColumnMig
}

// DropTableMig is returned after DropTable is called.
type DropTableMig interface {
	MigrationCallable
}

// AlterTableMig is returned after AlterTable is called.
type AlterTableMig interface {
	Rename(string) RenameMig
	AddColumn(string, string) AddColumnMig
	RenameColumn(string, string) RenameColumnMig
	DropColumn(string) DropColumnMig
	AddCons(string) AddConsMig
	DropPK(string) DropPKMig
	DropFK(string) DropFKMig
	DropUC(string) DropUCMig
}

// CreateIndexMig is returned after CreateIndex is called.
type CreateIndexMig interface {
	On(string, ...string) OnMig
}

// DropIndexMig is returned after DropIndex is called.
type DropIndexMig interface {
	MigrationCallable
}

// OnMig is returned after (*MigStmt).ON is called.
type OnMig interface {
	MigrationCallable
}

// RenameMig is returned after (*MigStmt).Rename is called.
type RenameMig interface {
	MigrationCallable
}

// AddColumnMig is returned after (*MigStmt).AddColumn is called.
type AddColumnMig interface {
	NotNull() NotNullMig
	AutoInc() AutoIncMig // Only MySQL
	Default(interface{}) DefaultMig
}

// DropClumnMig is returned after (*MigStmt).DropColumn is called.
type DropColumnMig interface {
	MigrationCallable
}

// RenameColumnMig is returned after (*MigStmt).RenameColumn is called.
type RenameColumnMig interface {
	MigrationCallable
}

// AddConsMig is returned after (*MigStmt).AddCons is called.
type AddConsMig interface {
	UC(...string) UCMig
	PK(...string) PKMig
	FK(...string) FKMig
}

// DropPKMig is returned after (*MigStmt).DropPK is called.
type DropPKMig interface {
	MigrationCallable
}

// DropFKMig is returned after (*MigStmt).DropFK is called.
type DropFKMig interface {
	MigrationCallable
}

// DropUCMig is returned after (*MigStmt).DropUC is called.
type DropUCMig interface {
	MigrationCallable
}

// ColumnMig is returned after (*MigStmt).Column is called.
type ColumnMig interface {
	Column(string, string) ColumnMig
	NotNull() NotNullMig
	AutoInc() AutoIncMig // Only MySQL
	Default(interface{}) DefaultMig
	Cons(string) ConsMig
	MigrationCallable
}

// NotNullMig is returned after (*MigStmt).NotNull is called.
type NotNullMig interface {
	Column(string, string) ColumnMig
	AutoInc() AutoIncMig
	Default(interface{}) DefaultMig
	Cons(string) ConsMig
	MigrationCallable
}

// AutoIncMig is returned after (*MigStmt).AutoInc is called.
type AutoIncMig interface {
	Column(string, string) ColumnMig
	Cons(string) ConsMig
	MigrationCallable
}

// DefaultMig is returned after (*MigStmt).Default is called.
type DefaultMig interface {
	Column(string, string) ColumnMig
	Cons(string) ConsMig
	MigrationCallable
}

// ConsMig is returned after (*MigStmt).Cons is called.
type ConsMig interface {
	UC(...string) UCMig
	PK(...string) PKMig
	FK(...string) FKMig
}

// UCMig is returned after (*MigStmt).UC is called.
type UCMig interface {
	Cons(string) ConsMig
	MigrationCallable
}

// PKMig is returned after (*MigStmt).PK is called.
type PKMig interface {
	Cons(string) ConsMig
	MigrationCallable
}

// FKMig is returned after (*MigStmt).FK is called.
type FKMig interface {
	Ref(string, string) RefMig
}

// RefMig is returned after (*MigStmt).Ref is called.
type RefMig interface {
	Cons(string) ConsMig
	MigrationCallable
}
