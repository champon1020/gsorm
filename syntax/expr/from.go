package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// From expression.
type From struct {
	Tables []syntax.Table
}

// Name returns string of clause.
func (f *From) Name() string {
	return "FROM"
}

func (f *From) addTable(col string) {
	c := syntax.NewTable(col)
	f.Tables = append(f.Tables, *c)
}

// String returns string of function call.
func (f *From) String() string {
	var s string
	for i, t := range f.Tables {
		if i != 0 {
			s += ", "
		}
		s += fmt.Sprintf("%q", t.Build())
	}
	return fmt.Sprintf("%s(%s)", f.Name(), s)
}

// Build make from statement set.
func (f *From) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(f.Name())
	for i, t := range f.Tables {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(t.Build())
	}
	return ss, nil
}

// NewFrom make new from object.
func NewFrom(tables []string) *From {
	f := new(From)
	for _, t := range tables {
		f.addTable(t)
	}
	return f
}
