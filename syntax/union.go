package syntax

import "fmt"

// Union expression.
type Union struct {
	Stmt Var
	All  bool
}

func (u *Union) name() string {
	n := "UNION"
	if u.All {
		n += " ALL"
	}
	return n
}

// String returns string of function call.
func (u *Union) String() string {
	return fmt.Sprintf("%s(%q)", u.name(), u.Stmt)
}

// Build make UNION statement set.
func (u *Union) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(u.name())
	ss.WriteValue(string(u.Stmt))
	return ss, nil
}

// NewUnion creates Union instance.
func NewUnion(stmt Var, all bool) *Union {
	return &Union{Stmt: stmt, All: all}
}
