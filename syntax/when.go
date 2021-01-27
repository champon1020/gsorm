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

// Then expression
type Then struct {
	Value interface{}
}

func (t *Then) name() string {
	return "THEN"
}

// String returns string of function call.
func (t *Then) String() string {
	switch v := t.Value.(type) {
	case string:
		return fmt.Sprintf("%s(%q)", t.name(), v)
	}
	return fmt.Sprintf("%s(%v)", t.name(), t.Value)
}

// Build makes THEN statement set.
func (t *Then) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(t.name())
	vStr, err := internal.ToString(t.Value)
	ss.WriteValue(vStr)
	return ss, err
}

// NewThen creates Then instance.
func NewThen(val interface{}) *Then {
	return &Then{Value: val}
}

// Else expression.
type Else struct {
	Value interface{}
}

func (e *Else) name() string {
	return "ELSE"
}

// String returns string of function call.
func (e *Else) String() string {
	switch v := e.Value.(type) {
	case string:
		return fmt.Sprintf("%s(%q)", e.name(), v)
	}
	return fmt.Sprintf("%s(%v)", e.name(), e.Value)
}

// Build makes ELSE statement set.
func (e *Else) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(e.name())
	vStr, err := internal.ToString(e.Value)
	ss.WriteValue(vStr)
	return ss, err
}

// NewElse creates Else instance.
func NewElse(val interface{}) *Else {
	return &Else{Value: val}
}
