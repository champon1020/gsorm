package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Cons is CONSTRAINT clause.
type Cons struct {
	Key string
}

// Keyword returns clause keyword.
func (c *Cons) Keyword() string {
	return "CONSTRAINT"
}

// Build makes CONSTRAINT clasue with syntax.StmtSet.
func (c *Cons) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Key)
	return ss, nil
}
