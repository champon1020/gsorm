package mig

import "github.com/champon1020/mgorm/syntax"

// Rename is RENAME TO clause.
type Rename struct {
	Table string
}

// Keyword returns clause keyword.
func (r *Rename) Keyword() string {
	return "RENAME TO"
}

// Build makes RENAME TO clause with syntax.StmtSet.
func (r *Rename) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(r.Table)
	return ss, nil
}
