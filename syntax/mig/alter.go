package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
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

func (a *AlterTable) String() string {
	return fmt.Sprintf("%s(%s)", a.Keyword(), a.Table)
}

// Build makes ALTER TABLE clause with syntax.StmtSet.
func (a *AlterTable) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Keyword())
	ss.WriteValue(a.Table)
	return ss, nil
}
