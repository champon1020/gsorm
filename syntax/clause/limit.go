package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Limit is LIMIT clause.
type Limit struct {
	Num int
}

// String returns function call as string.
func (l *Limit) String() string {
	return fmt.Sprintf("Limit(%v)", l.Num)
}

// Build creates the structure of LIMIT clause that implements interfaces.ClauseSet.
func (l *Limit) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("LIMIT")
	cs.WriteValue(strconv.Itoa(l.Num))
	return cs, nil
}
