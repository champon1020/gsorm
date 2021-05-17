package mig

import "github.com/champon1020/mgorm/syntax"

// DropColumn is DROP COLUMN clause.
type DropColumn struct {
	Column string
}

// Keyword returns clause keyword.
func (d *DropColumn) Keyword() string {
	return "DROP COLUMN"
}

// Build makes DROP COLUMN clause with syntax.StmtSet.
func (d *DropColumn) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Column)
	return ss, nil
}
