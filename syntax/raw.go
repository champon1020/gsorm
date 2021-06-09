package syntax

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
)

// RawClause is the raw string clause which can be defined by user.
type RawClause struct {
	RawStr string
	Values []interface{}
}

func (r *RawClause) String() string {
	s := fmt.Sprintf("%q", r.RawStr)
	if len(r.Values) > 0 {
		s += ", "
		s += internal.ToString(r.Values, nil)
	}
	return fmt.Sprintf("RAW CLAUSE(%s)", s)
}

// Build creates the pair of clause and value as ClauseSet.
func (r *RawClause) Build() (interfaces.ClauseSet, error) {
	s, err := BuildExpr(r.RawStr, r.Values...)
	if err != nil {
		return nil, err
	}
	ss := &ClauseSet{Keyword: s}
	return ss, nil
}
