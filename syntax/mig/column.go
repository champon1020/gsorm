package mig

import "github.com/champon1020/mgorm/syntax"

// Column is definition of table column.
type Column struct {
	Col  string
	Type string
}

// Name returns column name.
func (c *Column) Name() string {
	return c.Col
}

// Build makes database column definition with syntax.StmtSet.
func (c *Column) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Name())
	ss.WriteValue(c.Type)
	return ss, nil
}
