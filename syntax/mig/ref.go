package mig

import "github.com/champon1020/mgorm/syntax"

// Ref is REFERENCES clause.
type Ref struct {
	Table  string
	Column string
}

// Name returns clause keyword.
func (r *Ref) Name() string {
	return "REFERENCES"
}

// Build makes REFERENCES clause with syntax.StmtSet.
func (r *Ref) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Name())
	ss.WriteValue(r.Table)
	ss.WriteValue("(")
	ss.WriteValue(r.Column)
	ss.WriteValue(")")
	return ss, nil
}
