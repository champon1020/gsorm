package provider

// MigrationCallable is embedded into clause interfaces which can call (*MigStmt).Migration.
type MigrationCallable interface {
	Migration() error
	String() string
}
