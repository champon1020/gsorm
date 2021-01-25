package syntax

// Cmd is implemented by Select, Insert, Update and Delete types.
type Cmd interface {
	query() string
	Build() *StmtSet
}

// Expr is implemented by Where, And, and other types.
type Expr interface {
	name() string
	Build() (*StmtSet, error)
}

// Var type.
type Var string
