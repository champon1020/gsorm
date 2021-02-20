package mig

import "github.com/champon1020/mgorm/syntax"

// UC is UNIQUE clause.
type UC struct {
	Columns []string
}

// Keyword returns clause keyword.
func (u *UC) Keyword() string {
	return "UNIQUE"
}

// Build makes UNIQUE clause with syntax.StmtSet.
func (u *UC) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Keyword())
	ss.WriteValue("(")
	for i, c := range u.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	ss.WriteValue(")")
	return ss, nil
}
