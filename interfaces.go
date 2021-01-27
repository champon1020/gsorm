package mgorm

import "github.com/champon1020/mgorm/syntax"

// ExecutableStmt is executable interface of Stmt.
type ExecutableStmt interface {
	Query(interface{}) error
	ExpectQuery(interface{}) *Stmt
	Exec() error
	ExpectExec() *Stmt
	Var() syntax.Var
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
	Where(string, ...interface{}) WhereStmt
	GroupBy(...string) GroupByStmt
	OrderBy(string, bool) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// ValuesStmt is Stmt after Stmt.Values is executed.
type ValuesStmt interface {
	ExecutableStmt
}

// SetStmt is Stmt after Stmt.Set is executed.
type SetStmt interface {
	Where(string, ...interface{}) WhereStmt
	ExecutableStmt
}

// JoinStmt is Stmt after Stmt.Join, Stmt.LeftJoin, RightJoin or FullJoin is executed.
type JoinStmt interface {
	On(string, ...interface{}) OnStmt
}

// OnStmt is Stmt after Stmt.On is executed.
type OnStmt interface {
	Where(string, ...interface{}) WhereStmt
	OrderBy(string, bool) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// WhereStmt is Stmt after Stmt.Where is executed.
type WhereStmt interface {
	And(string, ...interface{}) AndStmt
	Or(string, ...interface{}) OrStmt
	GroupBy(...string) GroupByStmt
	OrderBy(string, bool) OrderByStmt
	Limit(int) LimitStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// AndStmt is Stmt after Stmt.And is executed.
type AndStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string, bool) OrderByStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// OrStmt is Stmt after Stmt.Or is executed.
type OrStmt interface {
	GroupBy(...string) GroupByStmt
	OrderBy(string, bool) OrderByStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// GroupByStmt is Stmt after Stmt.GroupBy is executed.
type GroupByStmt interface {
	Having(string, ...interface{}) HavingStmt
	OrderBy(string, bool) OrderByStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// HavingStmt is Stmt after Stmt.Having is executed.
type HavingStmt interface {
	OrderBy(string, bool) OrderByStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// OrderByStmt is Stmt after Stmt.OrderBy is executed.
type OrderByStmt interface {
	Limit(int) LimitStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// LimitStmt is Stmt after Stmt.Limit is executed.
type LimitStmt interface {
	Offset(int) OffsetStmt
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// OffsetStmt is Stmt after Stmt.Offset is executed.
type OffsetStmt interface {
	Union(syntax.Var) UnionStmt
	ExecutableStmt
}

// UnionStmt is Stmt after Stmt.Union is executed.
type UnionStmt interface {
	ExecutableStmt
}

// WhenStmt is Stmt after Stmt.When or mgorm.When is executed.
type WhenStmt interface {
	Then(interface{}) ThenStmt
}

// ThenStmt is Stmt after Stmt.Then is executed.
type ThenStmt interface {
	When(string, ...interface{}) WhenStmt
	Else(interface{}) ElseStmt
	Var() syntax.Var
}

// ElseStmt is Stmt after Stmt.Else is executed.
type ElseStmt interface {
	Var() syntax.Var
}
