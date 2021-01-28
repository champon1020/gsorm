package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
)

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
