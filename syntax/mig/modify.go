package mig

import "github.com/champon1020/mgorm/syntax"

// Modify is MODIFY clause.
type Modify struct {
	Column string
	Type   string
}

// Name returns clause keyword.
func (a *Modify) Name() string {
	return "MODIFY"
}

// Build makes MODIFY clause with syntax.StmtSet.
func (a *Modify) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Column)
	ss.WriteValue(a.Type)
	return ss, nil
}
