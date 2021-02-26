package xselect

import (
	"github.com/champon1020/mgorm/provider"
	"github.com/champon1020/mgorm/syntax"
)

// StmtMP is interface for returned value of mgorm.Select.
type StmtMP interface {
	From(...string) FromMP
}

// FromMP is interface for returned value of (*SelectStmt).From.
type FromMP interface {
	Join(string) JoinMP
	LeftJoin(string) JoinMP
	RightJoin(string) JoinMP
	FullJoin(string) JoinMP
	Where(string, ...interface{}) WhereMP
	GroupBy(...string) GroupByMP
	OrderBy(...string) OrderByMP
	Limit(int) LimitMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// JoinMP is interface for returned value of (*SelectStmt).Join.
type JoinMP interface {
	On(string, ...interface{}) OnMP
}

// OnMP is interface for returned value of (*SelectStmt).On.
type OnMP interface {
	Where(string, ...interface{}) WhereMP
	GroupBy(...string) GroupByMP
	OrderBy(...string) OrderByMP
	Limit(int) LimitMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// WhereMP is interface for returned value of (*SelectStmt).Where.
type WhereMP interface {
	And(string, ...interface{}) AndMP
	Or(string, ...interface{}) OrMP
	GroupBy(...string) GroupByMP
	OrderBy(...string) OrderByMP
	Limit(int) LimitMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// AndMP is interface for returned value of (*SelectStmt).And.
type AndMP interface {
	GroupBy(...string) GroupByMP
	OrderBy(...string) OrderByMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// OrMP is interface for returned value of (*SelectStmt).Or.
type OrMP interface {
	GroupBy(...string) GroupByMP
	OrderBy(...string) OrderByMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// GroupByMP is interface for returned value of (*SelectStmt).GroupBy.
type GroupByMP interface {
	Having(string, ...interface{}) HavingMP
	OrderBy(...string) OrderByMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// HavingMP is interface for returned value of (*SelectStmt).Having.
type HavingMP interface {
	OrderBy(...string) OrderByMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// OrderByMP is interface for returned value of (*SelectStmt).OrderBy.
type OrderByMP interface {
	Limit(int) LimitMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// LimitMP is interface for returned value of (*SelectStmt).Limit.
type LimitMP interface {
	Offset(int) OffsetMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// OffsetMP is interface for returned value of (*SelectStmt).Offset.
type OffsetMP interface {
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}

// UnionMP is interface for returned value of (*SelectStmt).Union.
type UnionMP interface {
	OrderBy(...string) OrderByMP
	Limit(int) LimitMP
	Union(syntax.Stmt) UnionMP
	UnionAll(syntax.Stmt) UnionMP
	provider.QueryCallable
}
