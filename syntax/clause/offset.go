package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Offset is OFFSET clause.
type Offset struct {
	Num int
}

// Keyword returns clause keyword.
func (o *Offset) Keyword() string {
	return "OFFSET"
}

// String returns function call with string.
func (o *Offset) String() string {
	return fmt.Sprintf("%s(%v)", o.Keyword(), o.Num)
}

// Build makes OFFSET clause with sytnax.StmtSet.
func (o *Offset) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Keyword())
	ss.WriteValue(strconv.Itoa(o.Num))
	return ss, nil
}
