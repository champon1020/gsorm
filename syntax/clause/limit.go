package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Limit is LIMIT clause.
type Limit struct {
	Num int
}

// String returns function call with string.
func (l *Limit) String() string {
	return fmt.Sprintf("Limit(%v)", l.Num)
}

// Build makes LIMIT clause with syntax.StmtSet.
func (l *Limit) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("LIMIT")
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}
