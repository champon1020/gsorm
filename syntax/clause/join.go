package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
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

// AddTable assigns the table to Join.Table.
func (j *Join) AddTable(table string) {
	j.Table = *syntax.NewTable(table)
}

// String returns function call as string.
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

// Build creates the structure of JOIN clause that implements interfaces.ClauseSet.
func (j *Join) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword(string(j.Type))
	cs.WriteValue(j.Table.Build())
	return cs, nil
}
