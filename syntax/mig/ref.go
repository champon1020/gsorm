package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Ref is REFERENCES clause.
type Ref struct {
	Table   string
	Columns []string
}

// Keyword returns clause keyword.
func (r *Ref) Keyword() string {
	return "REFERENCES"
}

func (r *Ref) String() string {
	return fmt.Sprintf("%s(%s, %v)", r.Keyword(), r.Table, r.Columns)
}

// Build makes REFERENCES clause with syntax.StmtSet.
func (r *Ref) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(r.Table)
	ss.WriteValue("(")
	for i, c := range r.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	ss.WriteValue(")")
	return ss, nil
}
