package syntax

import (
	"fmt"
	"strconv"
)

// Limit expression.
type Limit struct {
	Num int
}

func (l *Limit) name() string {
	return "LIMIT"
}

// String returns string of function call.
func (l *Limit) String() string {
	return fmt.Sprintf("%s(%v)", l.name(), l.Num)
}

// Build make limit statement set.
func (l *Limit) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(l.name())
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}

// NewLimit create new limit object.
func NewLimit(num int) *Limit {
	return &Limit{Num: num}
}
