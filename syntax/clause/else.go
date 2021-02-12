package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Else is ELSE clause.
type Else struct {
	Value    interface{}
	IsColumn bool
}

// Name returns clause keyword.
func (e *Else) Name() string {
	return "ELSE"
}

// String returns function call with string.
func (e *Else) String() string {
	switch v := e.Value.(type) {
	case string:
		return fmt.Sprintf("%s(%q)", e.Name(), v)
	}
	return fmt.Sprintf("%s(%v)", e.Name(), e.Value)
}

// Build makes ELSE clause with syntax.StmtSet.
func (e *Else) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(e.Name())
	vStr, err := internal.ToString(e.Value, !e.IsColumn)
	if err != nil {
		return nil, err
	}
	ss.WriteValue(vStr)
	return ss, nil
}

// NewElse creates Else instance.
func NewElse(val interface{}) *Else {
	return &Else{Value: val}
}
