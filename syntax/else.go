package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

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
