package mig

import "github.com/champon1020/mgorm/syntax"

// CreateTable is CREATE TABLE clause.
type CreateTable struct {
	Table string
}

// Keyword returns clause keyword.
func (c *CreateTable) Keyword() string {
	return "CREATE TABLE"
}

// Build makes CREATE TABLE clause with syntax.StmtSet.
func (c *CreateTable) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Table)
	return ss, nil
}
