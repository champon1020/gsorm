package mig

import "github.com/champon1020/mgorm/syntax"

// AddCons is ADD CONSTRAINT clause.
type AddCons struct {
	Key string
}

// Keyword returns clause keyword.
func (a *AddCons) Keyword() string {
	return "ADD CONSTRAINT"
}

// Build makes ADD CONSTRAINT clause with syntax.StmtSet.
func (a *AddCons) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Keyword())
	ss.WriteValue(a.Key)
	return ss, nil
}
