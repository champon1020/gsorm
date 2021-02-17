package mig

import "github.com/champon1020/mgorm/syntax"

// Add is ADD clause.
type Add struct {
	Column string
	Type   string
}

// Name returns clause keyword.
func (a *Add) Name() string {
	return "ADD"
}

// Build makes ADD clause with syntax.StmtSet.
func (a *Add) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Column)
	ss.WriteValue(a.Type)
	return ss, nil
}
