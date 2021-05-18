package domain

// Clause is interface for SQL clauses.
type Clause interface {
	Keyword() string
	String() string
	Build() (StmtSet, error)
}

// StmtSet is interface for the pair of clause and value.
type StmtSet interface {
	Build() string
	BuildValue() string
}
