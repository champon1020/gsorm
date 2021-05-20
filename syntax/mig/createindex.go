package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// CreateIndex is CREATE INDEX clause.
type CreateIndex struct {
	IdxName string
}

// Keyword returns clause keyword.
func (c *CreateIndex) Keyword() string {
	return "CREATE INDEX"
}

func (c *CreateIndex) String() string {
	return fmt.Sprintf("%s(%s)", c.Keyword(), c.IdxName)
}

// Build makes CREATE INDEX clause with syntax.StmtSet.
func (c *CreateIndex) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.IdxName)
	return ss, nil
}
