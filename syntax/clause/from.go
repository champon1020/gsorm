package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// From is FROM clause.
type From struct {
	Tables []syntax.Table
}

// AddTable appends the table to From.Tables.
func (f *From) AddTable(table string) {
	t := syntax.NewTable(table)
	f.Tables = append(f.Tables, *t)
}

// String returns function call as string.
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

// Build creates the structure of FROM clause that implements interfaces.ClauseSet.
func (f *From) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("FROM")
	for i, t := range f.Tables {
		if i != 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(t.Build())
	}
	return cs, nil
}
