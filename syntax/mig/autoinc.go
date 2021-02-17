package mig

import "github.com/champon1020/mgorm/syntax"

// AutoInc is AUTO_INCREMENT clause.
type AutoInc struct{}

// Name returns clause keyword.
func (a *AutoInc) Name() string {
	return "AUTO_INCREMENT"
}

// Build makes AUTO_INCREMENT clause with syntax.StmtSet.
func (a *AutoInc) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	return ss, nil
}
