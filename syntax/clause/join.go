package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// JoinType is type of JOIN clause.
type JoinType string

// Types of JOIN clause.
const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	FullJoin  JoinType = "FULL OUTER JOIN"
)

// Join is JOIN clause.
type Join struct {
	Table syntax.Table
	Type  JoinType
}

// Keyword returns clause keyword.
func (j *Join) Keyword() string {
	return string(j.Type)
}

// AddTable appends table to Join.
func (j *Join) AddTable(table string) {
	j.Table = *syntax.NewTable(table)
}

// String returns function call with string.
func (j *Join) String() string {
	return fmt.Sprintf("%s(%q)", j.Keyword(), j.Table.Build())
}

// Build makes JOIN clause with syntax.StmtSet.
func (j *Join) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(j.Keyword())
	ss.WriteValue(j.Table.Build())
	return ss, nil
}
