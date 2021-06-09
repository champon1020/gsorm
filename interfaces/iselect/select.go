package iselect

import (
	"github.com/champon1020/gsorm/interfaces"
)

// Stmt is interface which is returned by gsorm.Select.
type Stmt interface {
	RawClause(raw string, values ...interface{}) RawClause
	From(tables ...string) From
}

// RawClause is interface which is returned by (*Stmt).RawClause.
type RawClause interface {
	RawClause(raw string, values ...interface{}) RawClause
	From(tables ...string) From
	Join(table string) Join
	LeftJoin(table string) Join
	RightJoin(table string) Join
	On(expr string, values ...interface{}) On
	Where(expr string, values ...interface{}) Where
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	GroupBy(columns ...string) GroupBy
	Having(expr string, values ...interface{}) Having
	Union(stmt interfaces.Stmt) Union
	UnionAll(stmt interfaces.Stmt) Union
	OrderBy(columns ...string) OrderBy
	Limit(limit int) Limit
	Offset(offset int) Offset
	interfaces.QueryCallable
}

// From is interface which is returned by (*SelectStmt).From.
type From interface {
	RawClause(raw string, values ...interface{}) RawClause
	Join(table string) Join
	LeftJoin(table string) Join
	RightJoin(table string) Join
	On
	interfaces.QueryCallable
}

// Join is interface which is returned by (*SelectStmt).Join.
type Join interface {
	RawClause(raw string, values ...interface{}) RawClause
	On(expr string, values ...interface{}) On
}

// On is interface which is returned by (*SelectStmt).On.
type On interface {
	RawClause(raw string, values ...interface{}) RawClause
	Join(table string) Join
	LeftJoin(table string) Join
	RightJoin(table string) Join
	Where(expr string, values ...interface{}) Where
	Where
	interfaces.QueryCallable
}

// Where is interface which is returned by (*SelectStmt).Where.
type Where interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	Or(expr string, values ...interface{}) Or
	And
	interfaces.QueryCallable
}

// And is interface which is returned by (*SelectStmt).And.
type And interface {
	RawClause(raw string, values ...interface{}) RawClause
	And(expr string, values ...interface{}) And
	GroupBy(columns ...string) GroupBy
	GroupBy
	interfaces.QueryCallable
}

// Or is interface which is returned by (*SelectStmt).Or.
type Or interface {
	RawClause(raw string, values ...interface{}) RawClause
	Or(expr string, values ...interface{}) Or
	GroupBy(columns ...string) GroupBy
	GroupBy
	interfaces.QueryCallable
}

// GroupBy is interface which is returned by (*SelectStmt).GroupBy.
type GroupBy interface {
	RawClause(raw string, values ...interface{}) RawClause
	Having(expr string, values ...interface{}) Having
	Having
	interfaces.QueryCallable
}

// Having is interface which is returned by (*SelectStmt).Having.
type Having interface {
	RawClause(raw string, values ...interface{}) RawClause
	Union(stmt interfaces.Stmt) Union
	UnionAll(stmt interfaces.Stmt) Union
	Union
	interfaces.QueryCallable
}

// Union is interface which is returned by (*SelectStmt).Union.
type Union interface {
	RawClause(raw string, values ...interface{}) RawClause
	Union(stmt interfaces.Stmt) Union
	UnionAll(stmt interfaces.Stmt) Union
	OrderBy(columns ...string) OrderBy
	OrderBy
	interfaces.QueryCallable
}

// OrderBy is interface which is returned by (*SelectStmt).OrderBy.
type OrderBy interface {
	RawClause(raw string, values ...interface{}) RawClause
	OrderBy(columns ...string) OrderBy
	Limit(limit int) Limit
	interfaces.QueryCallable
}

// Limit is interface which is returned by (*SelectStmt).Limit.
type Limit interface {
	RawClause(raw string, values ...interface{}) RawClause
	Offset(int) Offset
	interfaces.QueryCallable
}

// Offset is interface which is returned by (*SelectStmt).Offset.
type Offset interface {
	RawClause(raw string, values ...interface{}) RawClause
	interfaces.QueryCallable
}
