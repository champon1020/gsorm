package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Ref is REFERENCES clause.
type Ref struct {
	Table  string
	Column string
}

// Keyword returns clause keyword.
func (r *Ref) Keyword() string {
	return "REFERENCES"
}

// Build makes REFERENCES clause with syntax.StmtSet.
func (r *Ref) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(fmt.Sprintf("%s(%s)", r.Table, r.Column))
	return ss, nil
}
