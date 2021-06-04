package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Delete is DELETE clause.
type Delete struct{}

// String returns function call with string.
func (d *Delete) String() string {
	return fmt.Sprintf("Delete()")
}

// Build makes DELETE clause with syntax.StmtSet.
func (d *Delete) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("DELETE")
	return ss, nil
}
