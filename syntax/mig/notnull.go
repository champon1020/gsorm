package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// NotNull is NOT NULL clause.
type NotNull struct{}

// Keyword returns clause keyword.
func (n *NotNull) Keyword() string {
	return "NOT NULL"
}

// Build makes NOT NULL clause with syntax.StmtSet.
func (n *NotNull) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(n.Keyword())
	return ss, nil
}
