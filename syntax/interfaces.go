package syntax

// Cmd interface.
type Cmd interface {
	query() string
	Build() *StmtSet
}

// Expr interface.
type Expr interface {
	name() string
	Build() (*StmtSet, error)
}
