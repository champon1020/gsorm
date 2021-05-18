package iselect

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/interfaces"
)

// Stmt is interface which is returned by mgorm.Select.
type Stmt interface {
	From(tables ...string) From
}

// From is interface which is returned by (*SelectStmt).From.
type From interface {
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
	On(string, ...interface{}) On
}

// On is interface which is returned by (*SelectStmt).On.
type On interface {
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
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// Having is interface which is returned by (*SelectStmt).Having.
type Having interface {
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}

// OrderBy is interface which is returned by (*SelectStmt).OrderBy.
type OrderBy interface {
	OrderBy(...string) OrderBy
	Limit(int) Limit
	interfaces.QueryCallable
}

// Limit is interface which is returned by (*SelectStmt).Limit.
type Limit interface {
	Offset(int) Offset
	interfaces.QueryCallable
}

// Offset is interface which is returned by (*SelectStmt).Offset.
type Offset interface {
	interfaces.QueryCallable
}

// Union is interface which is returned by (*SelectStmt).Union.
type Union interface {
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(domain.Stmt) Union
	UnionAll(domain.Stmt) Union
	interfaces.QueryCallable
}
