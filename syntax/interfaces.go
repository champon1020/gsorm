package syntax

// Cmd is implemented by Select, Insert, Update and Delete type.
type Cmd interface {
	query() string
	Build() *StmtSet
}

// Expr is implemented by Where, And, and other structures.
type Expr interface {
	name() string
	Build() (*StmtSet, error)
}
