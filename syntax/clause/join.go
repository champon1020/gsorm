package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// JoinType is type of JOIN clause.
type JoinType string

// Types of JOIN clause.
const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
)

// Join is JOIN clause.
type Join struct {
	Table syntax.Table
	Type  JoinType
}

// AddTable appends table to Join.
func (j *Join) AddTable(table string) {
	j.Table = *syntax.NewTable(table)
}

// String returns function call with string.
func (j *Join) String() string {
	var keyword string
	if j.Type == InnerJoin {
		keyword = "Join"
	}
	if j.Type == LeftJoin {
		keyword = "LeftJoin"
	}
	if j.Type == RightJoin {
		keyword = "RightJoin"
	}
	return fmt.Sprintf("%s(%q)", keyword, j.Table.Build())
}

// Build makes JOIN clause with syntax.StmtSet.
func (j *Join) Build() (domain.StmtSet, error) {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword(string(j.Type))
	ss.WriteValue(j.Table.Build())
	return ss, nil
}
