package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Update statement.
type Update struct {
	Table   syntax.Table
	Columns []string
}

// Query returns string of clause.
func (u *Update) Query() string {
	return "UPDATE"
}

func (u *Update) addTable(table string) {
	u.Table = *syntax.NewTable(table)
}

func (u *Update) addColumns(cols []string) {
	u.Columns = cols
}

// String returns string of function call.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	for _, c := range u.Columns {
		s += fmt.Sprintf(", %q", c)
	}
	return fmt.Sprintf("%s(%s)", u.Query(), s)
}

// Build make update statement set.
func (u *Update) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Query())
	ss.WriteValue(u.Table.Build())
	return ss
}

// NewUpdate create new update object.
func NewUpdate(table string, cols []string) *Update {
	u := new(Update)
	u.addTable(table)
	u.addColumns(cols)
	return u
}
