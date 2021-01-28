package expr

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/syntax"
)

// Limit expression.
type Limit struct {
	Num int
}

// Name returns string of clause.
func (l *Limit) Name() string {
	return "LIMIT"
}

// String returns string of function call.
func (l *Limit) String() string {
	return fmt.Sprintf("%s(%v)", l.Name(), l.Num)
}

// Build make limit statement set.
func (l *Limit) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(l.Name())
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}

// NewLimit create new limit object.
func NewLimit(num int) *Limit {
	return &Limit{Num: num}
}
