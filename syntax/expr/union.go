package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Union expression.
type Union struct {
	Stmt syntax.Var
	All  bool
}

// Name returns string of clause.
func (u *Union) Name() string {
	n := "UNION"
	if u.All {
		n += " ALL"
	}
	return n
}

// String returns string of function call.
func (u *Union) String() string {
	return fmt.Sprintf("%s(%q)", u.Name(), u.Stmt)
}

// Build make UNION statement set.
func (u *Union) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(u.Name())
	ss.WriteValue(string(u.Stmt))
	return ss, nil
}

// NewUnion creates Union instance.
func NewUnion(stmt syntax.Var, all bool) *Union {
	return &Union{Stmt: stmt, All: all}
}
