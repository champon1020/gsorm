package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// RenameColumn is RENAME COLUMN clause.
type RenameColumn struct {
	Column string
	Dest   string
}

// Keyword returns clause keyword.
func (r *RenameColumn) Keyword() string {
	return "RENAME COLUMN"
}

func (r *RenameColumn) String() string {
	return fmt.Sprintf("%s(%s, %s)", r.Keyword(), r.Column, r.Dest)
}

// Build makes RENAME COLUMN clause with syntax.StmtSet.
func (r *RenameColumn) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(r.Column)
	ss.WriteValue("TO")
	ss.WriteValue(r.Dest)
	return ss, nil
}
