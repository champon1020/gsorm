package xselect

import (
	"github.com/champon1020/mgorm/interfaces"
	"github.com/champon1020/mgorm/syntax"
)

// Stmt is interface which is returned by mgorm.Select.
type Stmt interface {
	From(...string) From
}

// From is interface which is returned by (*SelectStmt).From.
type From interface {
	Join(string) Join
	LeftJoin(string) Join
	RightJoin(string) Join
	FullJoin(string) Join
	Where(string, ...interface{}) Where
	GroupBy(...string) GroupBy
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Join is interface which is returned by (*SelectStmt).Join.
type Join interface {
	On(string, ...interface{}) On
}

// On is interface which is returned by (*SelectStmt).On.
type On interface {
	Where(string, ...interface{}) Where
	GroupBy(...string) GroupBy
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Where is interface which is returned by (*SelectStmt).Where.
type Where interface {
	And(string, ...interface{}) And
	Or(string, ...interface{}) Or
	GroupBy(...string) GroupBy
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// And is interface which is returned by (*SelectStmt).And.
type And interface {
	GroupBy(...string) GroupBy
	OrderBy(...string) OrderBy
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Or is interface which is returned by (*SelectStmt).Or.
type Or interface {
	GroupBy(...string) GroupBy
	OrderBy(...string) OrderBy
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// GroupBy is interface which is returned by (*SelectStmt).GroupBy.
type GroupBy interface {
	Having(string, ...interface{}) Having
	OrderBy(...string) OrderBy
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Having is interface which is returned by (*SelectStmt).Having.
type Having interface {
	OrderBy(...string) OrderBy
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// OrderBy is interface which is returned by (*SelectStmt).OrderBy.
type OrderBy interface {
	Limit(int) Limit
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Limit is interface which is returned by (*SelectStmt).Limit.
type Limit interface {
	Offset(int) Offset
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Offset is interface which is returned by (*SelectStmt).Offset.
type Offset interface {
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}

// Union is interface which is returned by (*SelectStmt).Union.
type Union interface {
	OrderBy(...string) OrderBy
	Limit(int) Limit
	Union(syntax.Stmt) Union
	UnionAll(syntax.Stmt) Union
	interfaces.QueryCallable
}
