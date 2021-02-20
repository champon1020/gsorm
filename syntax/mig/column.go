package mig

import "github.com/champon1020/mgorm/syntax"

// Column is definition of table column.
type Column struct {
	Col  string
	Type string
}

// Keyword returns column name.
func (c *Column) Keyword() string {
	return c.Col
}

// Build makes database column definition with syntax.StmtSet.
func (c *Column) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Type)
	return ss, nil
}
