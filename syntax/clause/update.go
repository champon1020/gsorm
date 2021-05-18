package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Update is UPDATE clause.
type Update struct {
	Table syntax.Table
}

// Keyword returns clause keyword.
func (u *Update) Keyword() string {
	return "UPDATE"
}

// AddTable appends table to Update.
func (u *Update) AddTable(table string) {
	u.Table = *syntax.NewTable(table)
}

// String returns function call with string.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	return fmt.Sprintf("%s(%s)", u.Keyword(), s)
}

// Build makes UPDATE clause with syntax.StmtSet.
func (u *Update) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(u.Keyword())
	ss.WriteValue(u.Table.Build())
	return ss, nil
}
