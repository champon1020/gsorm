package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Else expression.
type Else struct {
	Value interface{}
}

// Name returns string of clause.
func (e *Else) Name() string {
	return "ELSE"
}

// String returns string of function call.
func (e *Else) String() string {
	switch v := e.Value.(type) {
	case string:
		return fmt.Sprintf("%s(%q)", e.Name(), v)
	}
	return fmt.Sprintf("%s(%v)", e.Name(), e.Value)
}

// Build makes ELSE statement set.
func (e *Else) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(e.Name())
	vStr, err := internal.ToString(e.Value)
	ss.WriteValue(vStr)
	return ss, err
}

// NewElse creates Else instance.
func NewElse(val interface{}) *Else {
	return &Else{Value: val}
}
