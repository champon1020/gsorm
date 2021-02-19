package mig

import (
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

// Build makes DEFAULT clause with syntax.StmtSet.
func (d *Default) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	vStr, err := internal.ToString(d.Value, true)
	if err != nil {
		return nil, err
	}
	ss.WriteValue(vStr)
	return ss, nil
}
