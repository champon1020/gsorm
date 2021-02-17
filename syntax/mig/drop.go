package mig

import "github.com/champon1020/mgorm/syntax"

// Drop is DROP clause.
type Drop struct {
	Column string
}

// Name returns clause keyword.
func (a *Drop) Name() string {
	return "DROP"
}

// Build makes DROP clause with syntax.StmtSet.
func (a *Drop) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Column)
	return ss, nil
}
