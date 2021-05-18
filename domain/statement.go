package domain

// Stmt is interface for DeleteStmt, InsertStmt, SelectStmt, and so on.
type Stmt interface {
	String() string
	FuncString() string
	Called() []Clause
	Cmd() Clause
	CompareWith(targetStmt Stmt) error
}
