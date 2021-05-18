package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// DropTable is DROP TABLE clause.
type DropTable struct {
	Table string
}

// Keyword returns clause keyword.
func (d *DropTable) Keyword() string {
	return "DROP TABLE"
}

// Build makes DROP TABLE clause with syntax.StmtSet.
func (d *DropTable) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Table)
	return ss, nil
}
