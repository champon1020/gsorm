package expr

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/syntax"
)

// Offset expression.
type Offset struct {
	Num int
}

// Name returns string of clause.
func (o *Offset) Name() string {
	return "OFFSET"
}

// String returns string of function call.
func (o *Offset) String() string {
	return fmt.Sprintf("%s(%v)", o.Name(), o.Num)
}

// Build make offset statement set.
func (o *Offset) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(o.Name())
	ss.WriteValue(strconv.Itoa(o.Num))
	return ss, nil
}

// NewOffset create new offset object.
func NewOffset(num int) *Offset {
	return &Offset{Num: num}
}
