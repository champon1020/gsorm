package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// DropColumn is DROP COLUMN clause.
type DropColumn struct {
	Column string
}

// Keyword returns clause keyword.
func (d *DropColumn) Keyword() string {
	return "DROP COLUMN"
}

func (d *DropColumn) String() string {
	return fmt.Sprintf("%s(%s)", d.Keyword(), d.Column)
}

// Build makes DROP COLUMN clause with syntax.StmtSet.
func (d *DropColumn) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Column)
	return ss, nil
}
