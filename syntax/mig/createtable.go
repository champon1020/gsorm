package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// CreateTable is CREATE TABLE clause.
type CreateTable struct {
	Table string
}

// Keyword returns clause keyword.
func (c *CreateTable) Keyword() string {
	return "CREATE TABLE"
}

func (c *CreateTable) String() string {
	return fmt.Sprintf("%s(%s)", c.Keyword(), c.Table)
}

// Build makes CREATE TABLE clause with syntax.StmtSet.
func (c *CreateTable) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Table)
	return ss, nil
}
