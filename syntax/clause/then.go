package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Then is THEN clause.
type Then struct {
	Value    interface{}
	IsColumn bool
}

// Name returns clause keyword.
func (t *Then) Name() string {
	return "THEN"
}

// String returns function call with string.
func (t *Then) String() string {
	switch v := t.Value.(type) {
	case string:
		return fmt.Sprintf("%s(%q)", t.Name(), v)
	}
	return fmt.Sprintf("%s(%v)", t.Name(), t.Value)
}

// Build makes THEN clause with sytnax.StmtSet.
func (t *Then) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(t.Name())
	vStr, err := internal.ToString(t.Value, !t.IsColumn)
	if err != nil {
		return nil, err
	}
	ss.WriteValue(vStr)
	return ss, nil
}
