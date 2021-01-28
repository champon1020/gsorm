package syntax

import "fmt"

// JoinType is type of JOIN statement.
type JoinType string

// Types of JOIN statement.
const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	FullJoin  JoinType = "FULL OUTER JOIN"
)

// Join expression.
type Join struct {
	Table Table
	Type  JoinType
}

func (j *Join) name() string {
	return string(j.Type)
}

func (j *Join) addTable(table string) {
	j.Table = *NewTable(table)
}

// String returns string of function call.
func (j *Join) String() string {
	return fmt.Sprintf("%s(%q)", j.name(), j.Table.Build())
}

// Build make JOIN statement set.
func (j *Join) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(j.name())
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
