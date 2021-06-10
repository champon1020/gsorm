package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Update is UPDATE clause.
type Update struct {
	Table syntax.Table
}

// AddTable appends the table to Update.Table.
func (u *Update) AddTable(table string) {
	u.Table = *syntax.NewTable(table)
}

// String returns function call as string.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	return fmt.Sprintf("Update(%s)", s)
}

// Build creates the structure of UPDATE clause that implements interfaces.ClauseSet.
func (u *Update) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("UPDATE")
	cs.WriteValue(u.Table.Build())
	return cs, nil
}
