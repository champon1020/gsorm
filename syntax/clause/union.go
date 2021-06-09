package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
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

// String returns function call as string.
func (u *Union) String() string {
	keyword := "Union"
	if u.All {
		keyword += "All"
	}
	return fmt.Sprintf("%s(%q)", keyword, u.Stmt.SQL())
}

// Build creates the structure of UNION clause that implements interfaces.ClauseSet.
func (u *Union) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword(u.Keyword())
	cs.WriteValue(fmt.Sprintf("(%s)", u.Stmt.SQL()))
	return cs, nil
}
