package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Update is UPDATE clause.
type Update struct {
	Table   syntax.Table
	Columns []string
}

// Name returns clause keyword.
func (u *Update) Name() string {
	return "UPDATE"
}

// addTable appends table to Update.
func (u *Update) addTable(table string) {
	u.Table = *syntax.NewTable(table)
}

// addColumns appends columns to Update.
func (u *Update) addColumns(cols []string) {
	u.Columns = cols
}

// String returns function call with string.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	for _, c := range u.Columns {
		s += fmt.Sprintf(", %q", c)
	}
	return fmt.Sprintf("%s(%s)", u.Name(), s)
}

// Build makes UPDATE clause with syntax.StmtSet.
func (u *Update) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Name())
	ss.WriteValue(u.Table.Build())
	return ss, nil
}

// NewUpdate create new update object.
func NewUpdate(table string, cols []string) *Update {
	u := new(Update)
	u.addTable(table)
	u.addColumns(cols)
	return u
}
