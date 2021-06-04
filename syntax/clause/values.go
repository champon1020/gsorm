package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Values is VALUES clause.
type Values struct {
	Values []interface{}
}

// AddValue appends values to Values.
func (v *Values) AddValue(val interface{}) {
	v.Values = append(v.Values, val)
}

// String returns function call with string.
func (v *Values) String() string {
	s := internal.ToString(v.Values, &internal.ToStringOpt{DoubleQuotes: true})
	return fmt.Sprintf("Values(%s)", s)
}

// Build makes VALUES clause with sytnax.StmtSet.
func (v *Values) Build() (domain.StmtSet, error) {
	ss := &syntax.StmtSet{}
	ss.WriteKeyword("VALUES")
	ss.WriteValue("(")
	for i, v := range v.Values {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(internal.ToString(v, nil))
	}
	ss.WriteValue(")")
	return ss, nil
}
