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

// And clause.
type And struct {
	Expr   string
	Values []interface{}
}

func (a *And) name() string {
	return "AND"
}

// String returns string of function call.
func (a *And) String() string {
	s := fmt.Sprintf("%q", a.Expr)
	if len(a.Values) > 0 {
		s += ", "
		s += internal.SliceToString(a.Values)
	}
	return fmt.Sprintf("%s(%s)", a.name(), s)
}

// Build make AND statement set.
func (a *And) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(a.Expr, a.Values...)
	ss.WriteClause(a.name())
	ss.Parens = true
	return ss, err
}

// NewAnd create AND clause object.
func NewAnd(expr string, vals ...interface{}) *And {
	return &And{Expr: expr, Values: vals}
}

// Or clause.
type Or struct {
	Expr   string
	Values []interface{}
}

func (o *Or) name() string {
	return "OR"
}

// String returns string of function call.
func (o *Or) String() string {
	s := fmt.Sprintf("%q", o.Expr)
	if len(o.Values) > 0 {
		s += ", "
		s += internal.SliceToString(o.Values)
	}
	return fmt.Sprintf("%s(%s)", o.name(), s)
}

// Build make OR statement set.
func (o *Or) Build() (*StmtSet, error) {
	ss, err := buildStmtSet(o.Expr, o.Values...)
	ss.WriteClause(o.name())
	ss.Parens = true
	return ss, err
}

// NewOr create OR clause object.
func NewOr(expr string, vals ...interface{}) *Or {
	return &Or{Expr: expr, Values: vals}
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
