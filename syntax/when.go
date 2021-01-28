package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

// When expression.
type When struct {
	Expr   string
	Values []interface{}
}

func (w *When) name() string {
	return "WHEN"
}

// String returns string of function call.
func (w *When) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += fmt.Sprintf(", %s", internal.SliceToString(w.Values))
	}
	return fmt.Sprintf("%s(%s)", w.name(), s)
}

// Build makes WHEN statement set.
func (w *When) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(w.Expr, w.Values...)
	ss.WriteClause(w.name())
	return ss, err
}

// NewWhen creates When instance.
func NewWhen(expr string, vals ...interface{}) *When {
	return &When{Expr: expr, Values: vals}
}
