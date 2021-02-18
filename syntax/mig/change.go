package mig

import "github.com/champon1020/mgorm/syntax"

// Change is CHANGE clause.
type Change struct {
	Column string
	Dest   string
	Type   string
}

// Name returns clause keyword.
func (a *Change) Name() string {
	return "CHANGE"
}

// Build makes CHANGE clause with syntax.StmtSet.
func (a *Change) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Column)
	ss.WriteValue(a.Dest)
	ss.WriteValue(a.Type)
	return ss, nil
}
