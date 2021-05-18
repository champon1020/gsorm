package mig

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Foreign is FOREIGN KEY clasue.
type Foreign struct {
	Columns []string
}

// Keyword returns clause keyword.
func (f *Foreign) Keyword() string {
	return "FOREIGN KEY"
}

// Build makes FOREIGN KEY clasue with syntax.StmtSet.
func (f *Foreign) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(f.Keyword())
	ss.WriteValue("(")
	for i, c := range f.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	ss.WriteValue(")")
	return ss, nil
}
