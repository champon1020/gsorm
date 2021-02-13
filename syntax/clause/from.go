package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// From is FROM clause.
type From struct {
	Tables []syntax.Table
}

// Name returns clause keyword.
func (f *From) Name() string {
	return "FROM"
}

// AddTable appends table to From.
func (f *From) AddTable(table string) {
	t := syntax.NewTable(table)
	f.Tables = append(f.Tables, *t)
}

// String returns function call with string.
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

// Build makes FROM clause with syntax.StmtSet.
func (f *From) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(f.Name())
	for i, t := range f.Tables {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(t.Build())
	}
	return ss, nil
}
