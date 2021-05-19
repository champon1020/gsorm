package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Default is DEFAULT clause.
type Default struct {
	Value interface{}
}

// Keyword returns clause keyword.
func (d *Default) Keyword() string {
	return "DEFAULT"
}

func (d *Default) String() string {
	return fmt.Sprintf("%s(%s)", d.Keyword(), d.Value)
}

// Build makes DEFAULT clause with syntax.StmtSet.
func (d *Default) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(internal.ToString(d.Value, nil))
	return ss, nil
}
