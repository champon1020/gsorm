package syntax

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
)

// RawClause is the raw string clause which can be defined by user.
type RawClause struct {
	RawStr string
	Values []interface{}
}

// Name returns the raw string.
func (r *RawClause) Name() string {
	return r.RawStr
}

func (r *RawClause) String() string {
	s := fmt.Sprintf("%q", r.RawStr)
	if len(r.Values) > 0 {
		s += ", "
		s += internal.ToString(r.Values, nil)
	}
	return fmt.Sprintf("RAW CLAUSE(%s)", s)
}

// Build creates the pair of clause and value as syntax.StmtSet.
func (r *RawClause) Build() (domain.StmtSet, error) {
	s, err := BuildExpr(r.RawStr, r.Values...)
	if err != nil {
		return nil, err
	}
	ss := &StmtSet{Keyword: s}
	return ss, nil
}
