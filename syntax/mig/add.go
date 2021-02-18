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

// AddCons is ADD CONSTRAINT clause.
type AddConstraint struct {
	Key string
}

// Name returns clause keyword.
func (a *AddConstraint) Name() string {
	return "ADD CONSTRAINT"
}

// Build makes ADD CONSTRAINT clause with syntax.StmtSet.
func (a *AddConstraint) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Key)
	return ss, nil
}
