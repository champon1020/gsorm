package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Column is definition of table column.
type Column struct {
	Col  string
	Type string
}

// Keyword returns column name.
func (c *Column) Keyword() string {
	return c.Col
}

func (c *Column) String() string {
	return fmt.Sprintf("COLUMN(%s, %s)", c.Col, c.Type)
}

// Build makes database column definition with syntax.StmtSet.
func (c *Column) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Type)
	return ss, nil
}
