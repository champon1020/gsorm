package iselect

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/interfaces"
)

// Stmt is interface which is returned by mgorm.Select.
type Stmt interface {
	RawClause(rs string, v ...interface{}) RawClause
	From(tables ...string) From
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(rs string, v ...interface{}) RawClause
	From(t ...string) From
	Join(t string) Join
	LeftJoin(t string) Join
	RightJoin(t string) Join
	On(e string, v ...interface{}) On
	Where(e string, v ...interface{}) Where
	And(e string, v ...interface{}) And
	Or(e string, v ...interface{}) Or
	GroupBy(c ...string) GroupBy
	Having(e string, v ...interface{}) Having
	OrderBy(c ...string) OrderBy
	Limit(l int) Limit
	Offset(o int) Offset
	Union(s domain.Stmt) Union
	UnionAll(s domain.Stmt) Union
	interfaces.QueryCallable
}

// From is interface which is returned by (*SelectStmt).From.
type From interface {
	RawClause(rs string, v ...interface{}) RawClause
	Join(table string) Join
	LeftJoin(table string) Join
	RightJoin(table string) Join
	Where(expr string, vals ...interface{}) Where
	GroupBy(cols ...string) GroupBy
	Having(expr string, vals ...interface{}) Having
	OrderBy(cols ...string) OrderBy
	Limit(limit int) Limit
	Union(selectStmt domain.Stmt) Union
	UnionAll(selectStmt domain.Stmt) Union
	interfaces.QueryCallable
}

// Join is interface which is returned by (*SelectStmt).Join.
type Join interface {
	RawClause(rs string, v ...interface{}) RawClause
	On(string, ...interface{}) On
}

// On is interface which is returned by (*SelectStmt).On.
type On interface {
	RawClause(rs string, v ...interface{}) RawClause
	Join(string) Join
	LeftJoin(string) Join
	RightJoin(string) Join
	Where(string, ...interface{}) Where
	GroupBy(...string) GroupBy
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// Where is interface which is returned by (*SelectStmt).Where.
type Where interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	GroupBy(...string) GroupBy
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// And is interface which is returned by (*SelectStmt).And.
type And interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	GroupBy(...string) GroupBy
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// Or is interface which is returned by (*SelectStmt).Or.
type Or interface {
	RawClause(rs string, v ...interface{}) RawClause
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	GroupBy(...string) GroupBy
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// GroupBy is interface which is returned by (*SelectStmt).GroupBy.
type GroupBy interface {
	RawClause(rs string, v ...interface{}) RawClause
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// Having is interface which is returned by (*SelectStmt).Having.
type Having interface {
	RawClause(rs string, v ...interface{}) RawClause
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// OrderBy is interface which is returned by (*SelectStmt).OrderBy.
type OrderBy interface {
	RawClause(rs string, v ...interface{}) RawClause
	OrderBy(...string) OrderBy
	Limit(int) Limit
	interfaces.QueryCallable
}

// Limit is interface which is returned by (*SelectStmt).Limit.
type Limit interface {
	RawClause(rs string, v ...interface{}) RawClause
	Offset(int) Offset
	interfaces.QueryCallable
}

// Offset is interface which is returned by (*SelectStmt).Offset.
type Offset interface {
	RawClause(rs string, v ...interface{}) RawClause
	interfaces.QueryCallable
}

// Union is interface which is returned by (*SelectStmt).Union.
type Union interface {
	RawClause(rs string, v ...interface{}) RawClause
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}
