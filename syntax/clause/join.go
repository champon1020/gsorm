package clause

import (
	"fmt"

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

// Name returns clause keyword.
func (j *Join) Name() string {
	return string(j.Type)
}

// addTable appends table to Join.
func (j *Join) addTable(table string) {
	j.Table = *syntax.NewTable(table)
}

// String returns function call with string.
func (j *Join) String() string {
	return fmt.Sprintf("%s(%q)", j.Name(), j.Table.Build())
}

// Build makes JOIN clause with syntax.StmtSet.
func (j *Join) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(j.Name())
	ss.WriteValue(j.Table.Build())
	return ss, nil
}

// NewJoin create Join instance.
func NewJoin(table string, typ JoinType) *Join {
	j := new(Join)
	j.Type = typ
	j.addTable(table)
	return j
}
