package mig

import "github.com/champon1020/mgorm/syntax"

// Rename is RENAME TO clause.
type Rename struct {
	Table string
}

// Name returns clause keyword.
func (r *Rename) Name() string {
	return "RENAME TO"
}

// Build makes RENAME TO clause with syntax.StmtSet.
func (r *Rename) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Name())
	ss.WriteValue(r.Table)
	return ss, nil
}
