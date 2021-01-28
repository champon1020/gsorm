package syntax

import (
	"fmt"
	"strconv"
)

// Offset expression.
type Offset struct {
	Num int
}

func (o *Offset) name() string {
	return "OFFSET"
}

// String returns string of function call.
func (o *Offset) String() string {
	return fmt.Sprintf("%s(%v)", o.name(), o.Num)
}

// Build make offset statement set.
func (o *Offset) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(o.name())
	ss.WriteValue(strconv.Itoa(o.Num))
	return ss, nil
}

// NewOffset create new offset object.
func NewOffset(num int) *Offset {
	return &Offset{Num: num}
}
