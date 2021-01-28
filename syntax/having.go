package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

// Having expression.
type Having struct {
	Expr   string
	Values []interface{}
}

func (h *Having) name() string {
	return "HAVING"
}

// String returns string of function call.
func (h *Having) String() string {
	s := fmt.Sprintf("%q", h.Expr)
	if len(h.Values) > 0 {
		s += ", "
		s += internal.SliceToString(h.Values)
	}
	return fmt.Sprintf("%s(%s)", h.name(), s)
}

// Build makes HAVING statement set.
func (h *Having) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(h.Expr, h.Values...)
	ss.WriteClause(h.name())
	return ss, err
}

// NewHaving creates Having instance.
func NewHaving(expr string, vals ...interface{}) *Having {
	return &Having{Expr: expr, Values: vals}
}
