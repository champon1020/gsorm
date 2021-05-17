package mig

import "github.com/champon1020/mgorm/syntax"

// CreateDB is CREATE DATABASE clause.
type CreateDB struct {
	DBName string
}

// Keyword returns clause keyword.
func (c *CreateDB) Keyword() string {
	return "CREATE DATABASE"
}

// Build makes CREATE DATABASE clause with syntax.StmtSet.
func (c *CreateDB) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.DBName)
	return ss, nil
}
