package mig

import "github.com/champon1020/mgorm/syntax"

// RenameColumn is RENAME COLUMN clause.
type RenameColumn struct {
	Column string
	Dest   string
}

// Keyword returns clause keyword.
func (r *RenameColumn) Keyword() string {
	return "RENAME COLUMN"
}

// Build makes RENAME COLUMN clause with syntax.StmtSet.
func (r *RenameColumn) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(r.Column)
	ss.WriteValue("TO")
	ss.WriteValue(r.Dest)
	return ss, nil
}
