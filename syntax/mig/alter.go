package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// AlterTable is ALTER TABLE clause.
type AlterTable struct {
	Table string
}

// Keyword returns clause keyword.
func (a *AlterTable) Keyword() string {
	return "ALTER TABLE"
}

// Build makes ALTER TABLE clause with syntax.StmtSet.
func (a *AlterTable) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Keyword())
	ss.WriteValue(a.Table)
	return ss, nil
}
