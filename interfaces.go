package mgorm

import (
	"github.com/champon1020/mgorm/syntax"
)

// sqlDB is interface that is implemented by *sql.DB.
type sqlDB interface {
	Ping() error
}

// QueryCallable is embedded into interfaces which is callable Stmt.Query.
type QueryCallable interface {
	Query(interface{}) error
	ExpectQuery(interface{}) *Stmt
	Sub() syntax.Sub
	String() string
}

// ExecCallable is embedded into interfaces which is callable Stmt.Exec.
type ExecCallable interface {
	Exec() error
	ExpectExec() *Stmt
	String() string
}

// SelectStmt is Stmt after Select is executed.
type SelectStmt interface {
	From(...string) FromStmt
}

// InsertStmt is Stmt after Insert is executed.
type InsertStmt interface {
	Values(...interface{}) ValuesStmt
}

// UpdateStmt is Stmt after Update is executed.
type UpdateStmt interface {
	Set(...interface{}) SetStmt
}

// DeleteStmt is Stmt after Delete is executed.
type DeleteStmt interface {
	From(...string) FromStmt
}

// FromStmt is Stmt after Stmt.From is executed.
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

// ValuesStmt is Stmt after Stmt.Values is executed.
type ValuesStmt interface {
	ExecCallable
}

// SetStmt is Stmt after Stmt.Set is executed.
type SetStmt interface {
	Where(string, ...interface{}) WhereStmt
	ExecCallable
}

// JoinStmt is Stmt after Stmt.Join, Stmt.LeftJoin, RightJoin or FullJoin is executed.
type JoinStmt interface {
	On(string, ...interface{}) OnStmt
}

// OnStmt is Stmt after Stmt.On is executed.
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

// WhereStmt is Stmt after Stmt.Where is executed.
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

// AndStmt is Stmt after Stmt.And is executed.
type AndStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
	ExecCallable
}

// OrStmt is Stmt after Stmt.Or is executed.
type OrStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
	ExecCallable
}

// GroupByStmt is Stmt after Stmt.GroupBy is executed.
type GroupByStmt interface {
	Having(string, ...interface{}) HavingStmt
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// HavingStmt is Stmt after Stmt.Having is executed.
type HavingStmt interface {
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// OrderByStmt is Stmt after Stmt.OrderBy is executed.
type OrderByStmt interface {
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// LimitStmt is Stmt after Stmt.Limit is executed.
type LimitStmt interface {
	Offset(int) OffsetStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// OffsetStmt is Stmt after Stmt.Offset is executed.
type OffsetStmt interface {
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// UnionStmt is Stmt after Stmt.Union is executed.
type UnionStmt interface {
	OrderBy(string) OrderByStmt
	OrderByDesc(string) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Sub) UnionStmt
	UnionAll(syntax.Sub) UnionStmt
	QueryCallable
}

// WhenStmt is Stmt after Stmt.When or mgorm.When is executed.
type WhenStmt interface {
	Then(interface{}) ThenStmt
}

// ThenStmt is Stmt after Stmt.Then is executed.
type ThenStmt interface {
	When(string, ...interface{}) WhenStmt
	Else(interface{}) ElseStmt
	CaseColumn() string
	CaseValue() string
}

// ElseStmt is Stmt after Stmt.Else is executed.
type ElseStmt interface {
	CaseColumn() string
	CaseValue() string
	QueryCallable
}