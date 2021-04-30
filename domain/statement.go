package domain

import "github.com/champon1020/mgorm/syntax"

// Stmt is interface for DeleteStmt, InsertStmt, SelectStmt, and so on.
type Stmt interface {
	String() string
	FuncString() string
	Called() []syntax.Clause
	Cmd() syntax.Clause
}
