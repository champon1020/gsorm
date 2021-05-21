package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// CreateDB is CREATE DATABASE clause.
type CreateDB struct {
	DBName string
}

// Keyword returns clause keyword.
func (c *CreateDB) Keyword() string {
	return "CREATE DATABASE"
}

func (c *CreateDB) String() string {
	return fmt.Sprintf("%s(%s)", c.Keyword(), c.DBName)
}

// Build makes CREATE DATABASE clause with syntax.StmtSet.
func (c *CreateDB) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.DBName)
	return ss, nil
}
