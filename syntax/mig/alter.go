package mig

import "github.com/champon1020/mgorm/syntax"

// AlterTable is ALTER TABLE clause.
type AlterTable struct {
	Table string
}

// Query returns clause keyword.
func (a *AlterTable) Query() string {
	return "ALTER TABLE"
}

// Build makes ALTER TABLE clause with syntax.StmtSet.
func (a *AlterTable) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Query())
	ss.WriteValue(a.Table)
	return ss
}
