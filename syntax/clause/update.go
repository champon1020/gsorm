package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Update is UPDATE clause.
type Update struct {
	Table syntax.Table
}

// AddTable appends table to Update.
func (u *Update) AddTable(table string) {
	u.Table = *syntax.NewTable(table)
}

// String returns function call with string.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	return fmt.Sprintf("Update(%s)", s)
}

// Build makes UPDATE clause with syntax.StmtSet.
func (u *Update) Build() (domain.StmtSet, error) {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword("UPDATE")
	ss.WriteValue(u.Table.Build())
	return ss, nil
}
