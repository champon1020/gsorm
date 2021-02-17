package mig

import "github.com/champon1020/mgorm/syntax"

// NotNull is NOT NULL clause.
type NotNull struct{}

// Name returns clause keyword.
func (n *NotNull) Name() string {
	return "NOT NULL"
}

// Build makes NOT NULL clause with syntax.StmtSet.
func (n *NotNull) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(n.Name())
	return ss, nil
}
