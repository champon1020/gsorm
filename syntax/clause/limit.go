package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Limit is LIMIT clause.
type Limit struct {
	Num int
}

// Keyword returns clause keyword.
func (l *Limit) Keyword() string {
	return "LIMIT"
}

// String returns function call with string.
func (l *Limit) String() string {
	return fmt.Sprintf("%s(%v)", l.Keyword(), l.Num)
}

// Build makes LIMIT clause with syntax.StmtSet.
func (l *Limit) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(l.Keyword())
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}
