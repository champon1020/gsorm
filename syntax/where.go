package syntax

import (
	"errors"
	"fmt"
	"strings"

	"github.com/champon1020/mgorm/internal"
)

// Op values for error handling.
const (
	OpBuildStmtSet internal.Op = "syntax.buildStmtSet"
)

// Where clause.
type Where struct {
	Expr   string
	Values []interface{}
}

func (w *Where) name() string {
	return "WHERE"
}

// String returns string of function call.
func (w *Where) String() string {
	s := fmt.Sprintf("%q", w.Expr)
	if len(w.Values) > 0 {
		s += ", "
		s += internal.SliceToString(w.Values)
	}
	return fmt.Sprintf("%s(%s)", w.name(), s)
}

// Build make WHERE statement set.
func (w *Where) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(w.Expr, w.Values...)
	ss.WriteClause(w.name())
	return ss, err
}

// NewWhere create WHERE clause object.
func NewWhere(expr string, vals ...interface{}) *Where {
	return &Where{Expr: expr, Values: vals}
}

// buildStmtSet make StmtSet with expr and values.
func buildStmtSet(expr string, vals ...interface{}) (*StmtSet, error) {
	if strings.Count(expr, "?") != len(vals) {
		err := errors.New("Length of values is not valid")
		return nil, internal.NewError(OpBuildStmtSet, internal.KindBasic, err)
	}

	ss := new(StmtSet)
	values := []interface{}{}
	for _, v := range vals {
		if sel, ok := v.(Var); ok {
			values = append(values, fmt.Sprintf("(%s)", sel))
			continue
		}
		vStr, err := internal.ToString(v)
		if err != nil {
			return nil, err
		}
		values = append(values, vStr)
	}

	ss.WriteValue(fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...))
	return ss, nil
}
