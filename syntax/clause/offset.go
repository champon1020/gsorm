package clause

import (
	"fmt"
	"strconv"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/syntax"
)

// Offset is OFFSET clause.
type Offset struct {
	Num int
}

// String returns function call with string.
func (o *Offset) String() string {
	return fmt.Sprintf("Offset(%v)", o.Num)
}

// Build makes OFFSET clause with sytnax.StmtSet.
func (o *Offset) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword("OFFSET")
	ss.WriteValue(strconv.Itoa(o.Num))
	return ss, nil
}
