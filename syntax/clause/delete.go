package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Delete is DELETE clause.
type Delete struct{}

// Keyword returns clause keyword.
func (d *Delete) Keyword() string {
	return "DELETE"
}

// String returns function call with string.
func (d *Delete) String() string {
	return fmt.Sprintf("%s()", d.Keyword())
}

// Build makes DELETE clause with syntax.StmtSet.
func (d *Delete) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	return ss, nil
}
