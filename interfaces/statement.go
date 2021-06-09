package interfaces

// Clause is interface for SQL clauses.
type Clause interface {
	String() string
	Build() (ClauseSet, error)
}

// ClauseSet is interface for the pair of clause and value.
type ClauseSet interface {
	Build() string
	BuildValue() string
}

// Stmt is the interface for statements.
type Stmt interface {
	SQL() string
	String() string
	Clauses() []Clause
	Cmd() Clause
	CompareWith(s Stmt) error
}
