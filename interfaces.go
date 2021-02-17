package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/syntax"
)

// Pool is database connection pool like DB or Tx. This is also implemented by MockDB and MockTx.
type Pool interface {
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
	Pool
	Ping() error
	SetConnMaxLifetime(n time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Close() error
	Begin() (*sql.Tx, error)
}

// sqlTx is interface for sql.Tx.
type sqlTx interface {
	Pool
	Commit() error
	Rollback() error
}

// QueryCallable is embedded into clause interfaces which can call (*Stmt).Query.
type QueryCallable interface {
	Query(interface{}) error
	ExpectQuery(interface{}) *Stmt
	Sub() syntax.Sub
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
}

// UpdateStmt is returned after Update is called.
type UpdateStmt interface {
	Set(...interface{}) SetStmt
}

// DeleteStmt is returned after Delete is called.
type DeleteStmt interface {
	From(...string) FromStmt
}

// FromStmt is returned after (*Stmt).From is called.
type FromStmt interface {
	Join(string) JoinStmt
	LeftJoin(string) JoinStmt
	RightJoin(string) JoinStmt
	FullJoin(string) JoinStmt
	Where(string, ...interface{}) WhereStmt
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
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
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// WhereStmt is returned after (*Stmt).Where is called.
type WhereStmt interface {
	And(string, ...interface{}) AndStmt
	Or(string, ...interface{}) OrStmt
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
	ExecCallable
}

// AndStmt is returned after (*Stmt).And is called.
type AndStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
	ExecCallable
}

// OrStmt is returned after (*Stmt).Or is called.
type OrStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
	ExecCallable
}

// GroupByStmt is returned after (*Stmt).GroupBy is called.
type GroupByStmt interface {
	Having(string, ...interface{}) HavingStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// HavingStmt is returned after (*Stmt).Having is called.
type HavingStmt interface {
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// OrderByStmt is returned after (*Stmt).OrderBy is called.
type OrderByStmt interface {
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// LimitStmt is returned after (*Stmt).Limit is called.
type LimitStmt interface {
	Offset(int) OffsetStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// OffsetStmt is returned after (*Stmt).Offset is called.
type OffsetStmt interface {
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// UnionStmt is returned after (*Stmt).Union is called.
type UnionStmt interface {
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// WhenStmt is returned after (*Stmt).When or mgorm.When is called.
type WhenStmt interface {
	Then(interface{}) ThenStmt
}

// ThenStmt is returned after (*Stmt).Then is called.
type ThenStmt interface {
	When(string, ...interface{}) WhenStmt
	Else(interface{}) ElseStmt
	CaseColumn() string
	CaseValue() string
}

// ElseStmt is returned after (*Stmt).Else is called.
type ElseStmt interface {
	CaseColumn() string
	CaseValue() string
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

// CreateTableMig is returned after CreateTable is called.
type CreateTableMig interface {
	Column(string, string) ColumnMig
}

// ColumnMig is returned after (*MigStmt).Column is called.
type ColumnMig interface {
	Column(string, string) ColumnMig
	NotNull() NotNullMig
	AutoInc() AutoIncMig
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
	PK(string) PKMig
	FK(string) FKMig
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
