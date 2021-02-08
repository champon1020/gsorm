package syntax

// Cmd is implemented by Select, Insert, Update and Delete types.
type Cmd interface {
	Query() string
	String() string
	Build() *StmtSet
}

// Expr is implemented by Where, And, and other types.
type Expr interface {
	Name() string
	String() string
	Build() (*StmtSet, error)
}

// Sub type.
type Sub string
