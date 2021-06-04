package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// From is FROM clause.
type From struct {
	Tables []syntax.Table
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
	return fmt.Sprintf("From(%s)", s)
}

// Build makes FROM clause with syntax.StmtSet.
func (f *From) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("FROM")
	for i, t := range f.Tables {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(t.Build())
	}
	return ss, nil
}
