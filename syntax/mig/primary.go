package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Primary is PRIMARY KEY clause.
type Primary struct {
	Columns []string
}

// Keyword returns clause keyword.
func (p *Primary) Keyword() string {
	return "PRIMARY KEY"
}

// Build makes PRIMARY KEY clause with syntax.StmtSet.
func (p *Primary) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(p.Keyword())
	ss.WriteValue("(")
	for i, c := range p.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	ss.WriteValue(")")
	return ss, nil
}
