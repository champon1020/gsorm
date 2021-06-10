package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Offset is OFFSET clause.
type Offset struct {
	Num int
}

// String returns function call as string.
func (o *Offset) String() string {
	return fmt.Sprintf("Offset(%v)", o.Num)
}

// Build creates the structure of OFFSET clause that implements interfaces.ClauseSet.
func (o *Offset) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("OFFSET")
	cs.WriteValue(strconv.Itoa(o.Num))
	return cs, nil
}
