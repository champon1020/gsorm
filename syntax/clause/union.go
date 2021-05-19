package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Union is UNION clause.
type Union struct {
	Stmt domain.Stmt
	All  bool
}

// Keyword returns clause keyword.
func (u *Union) Keyword() string {
	n := "UNION"
	if u.All {
		n += " ALL"
	}
	return n
}

// String returns function call with string.
func (u *Union) String() string {
	return fmt.Sprintf("%s(%q)", u.Keyword(), u.Stmt)
}

// Build makes UNION clause with syntax.StmtSet.
func (u *Union) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Keyword())
	ss.WriteValue(fmt.Sprintf("(%s)", u.Stmt.String()))
	return ss, nil
}
