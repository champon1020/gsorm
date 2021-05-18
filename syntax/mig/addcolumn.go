package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// AddColumn is ADD clause.
type AddColumn struct {
	Column string
	Type   string
}

// Keyword returns clause keyword.
func (a *AddColumn) Keyword() string {
	return "ADD COLUMN"
}

// Build makes ADD COLUMN clause with syntax.StmtSet.
func (a *AddColumn) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Keyword())
	ss.WriteValue(a.Column)
	ss.WriteValue(a.Type)
	return ss, nil
}
