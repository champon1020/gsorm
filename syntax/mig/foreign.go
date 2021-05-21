package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Foreign is FOREIGN KEY clasue.
type Foreign struct {
	Columns []string
}

// Keyword returns clause keyword.
func (f *Foreign) Keyword() string {
	return "FOREIGN KEY"
}

func (f *Foreign) String() string {
	return fmt.Sprintf("%s(%v)", f.Keyword(), f.Columns)
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
