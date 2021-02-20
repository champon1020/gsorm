package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Union is UNION clause.
type Union struct {
	Stmt syntax.Stmt
	All  bool
}

// Name returns clause keyword.
func (u *Union) Name() string {
	n := "UNION"
	if u.All {
		n += " ALL"
	}
	return n
}

// String returns function call with string.
func (u *Union) String() string {
	return fmt.Sprintf("%s(%q)", u.Name(), u.Stmt)
}

// Build makes UNION clause with syntax.StmtSet.
func (u *Union) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Name())
	ss.WriteValue(u.Stmt.String())
	return ss, nil
}
