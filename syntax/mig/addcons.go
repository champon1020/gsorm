package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// AddCons is ADD CONSTRAINT clause.
type AddCons struct {
	Key string
}

// Keyword returns clause keyword.
func (a *AddCons) Keyword() string {
	return "ADD CONSTRAINT"
}

func (a *AddCons) String() string {
	return fmt.Sprintf("%s(%s)", a.Keyword(), a.Key)
}

// Build makes ADD CONSTRAINT clause with syntax.StmtSet.
func (a *AddCons) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Keyword())
	ss.WriteValue(a.Key)
	return ss, nil
}
