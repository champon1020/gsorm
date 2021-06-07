package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Union is UNION clause.
type Union struct {
	Stmt interfaces.Stmt
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
	keyword := "Union"
	if u.All {
		keyword += "All"
	}
	return fmt.Sprintf("%s(%q)", keyword, u.Stmt.SQL())
}

// Build makes UNION clause with syntax.StmtSet.
func (u *Union) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Keyword())
	ss.WriteValue(fmt.Sprintf("(%s)", u.Stmt.SQL()))
	return ss, nil
}
