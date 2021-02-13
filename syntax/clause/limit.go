package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/syntax"
)

// Limit is LIMIT clause.
type Limit struct {
	Num int
}

// Name returns clause keyword.
func (l *Limit) Name() string {
	return "LIMIT"
}

// String returns function call with string.
func (l *Limit) String() string {
	return fmt.Sprintf("%s(%v)", l.Name(), l.Num)
}

// Build makes LIMIT clause with syntax.StmtSet.
func (l *Limit) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(l.Name())
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}
