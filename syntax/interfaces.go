package syntax

// Cmd is interface for entry commands like SELECT, INSERT, UPDATE or DELETE.
type Cmd interface {
	Query() string
	String() string
	Build() *StmtSet
}

// Clause is interface for SQL clauses.
type Clause interface {
	Name() string
	String() string
	Build() (*StmtSet, error)
}

// MigClause is interface for SQL clauses about database migration.
type MigClause interface {
	Keyword() string
	Build() (*StmtSet, error)
}

// Sub is type of subquery like ...WHERE (SELECT ...).
type Sub string
