package mig

import "github.com/champon1020/mgorm/syntax"

// PrimaryKey is PRIMARY KEY clause.
type PrimaryKey struct {
	Columns []string
}

// Name returns clause keyword.
func (p *PrimaryKey) Name() string {
	return "PRIMARY KEY"
}

// AddColumns appends columns to PrimaryKey.
func (p *PrimaryKey) AddColumns(col ...string) {
	p.Columns = append(p.Columns, col...)
}

// Build makes PRIMARY KEY clause with syntax.StmtSet.
func (p *PrimaryKey) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(p.Name())
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
