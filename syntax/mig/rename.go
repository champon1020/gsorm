package mig

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Rename is RENAME TO clause.
type Rename struct {
	Table string
}

// Keyword returns clause keyword.
func (r *Rename) Keyword() string {
	return "RENAME TO"
}

func (r *Rename) String() string {
	return fmt.Sprintf("%s(%s)", r.Keyword(), r.Table)
}

// Build makes RENAME TO clause with syntax.StmtSet.
func (r *Rename) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(r.Keyword())
	ss.WriteValue(r.Table)
	return ss, nil
}
