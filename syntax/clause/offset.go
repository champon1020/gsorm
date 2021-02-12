package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/syntax"
)

// Offset is OFFSET clause.
type Offset struct {
	Num int
}

// Name returns clause keyword.
func (o *Offset) Name() string {
	return "OFFSET"
}

// String returns function call with string.
func (o *Offset) String() string {
	return fmt.Sprintf("%s(%v)", o.Name(), o.Num)
}

// Build makes OFFSET clause with sytnax.StmtSet.
func (o *Offset) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Name())
	ss.WriteValue(strconv.Itoa(o.Num))
	return ss, nil
}

// NewOffset create new offset object.
func NewOffset(num int) *Offset {
	return &Offset{Num: num}
}
