package syntax

import "github.com/champon1020/mgorm/internal"

// When expression.
type When struct {
	Expr   string
	Values []interface{}
}

func (w *When) name() string {
	return "WHEN"
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
