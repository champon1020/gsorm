package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
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

func (d *DropTable) String() string {
	return fmt.Sprintf("%s(%s)", d.Keyword(), d.Table)
}

// Build makes DROP TABLE clause with syntax.StmtSet.
func (d *DropTable) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Table)
	return ss, nil
}
