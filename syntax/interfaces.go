package syntax

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

// Stmt is interface implementing mgorm.Stmt.
type Stmt interface {
	String() string
	Query(interface{}) error
}
