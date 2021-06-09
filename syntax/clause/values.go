package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
)

// Values is VALUES clause.
type Values struct {
	Values []interface{}
}

// AddValue appends the value to Values.Values.
func (v *Values) AddValue(val interface{}) {
	v.Values = append(v.Values, val)
}

// String returns function call as string.
func (v *Values) String() string {
	s := internal.ToString(v.Values, &internal.ToStringOpt{DoubleQuotes: true})
	return fmt.Sprintf("Values(%s)", s)
}

// Build creates the structure of VALUES clause that implements interfaces.ClauseSet.
func (v *Values) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("VALUES")
	cs.WriteValue("(")
	for i, v := range v.Values {
		if i != 0 {
			cs.WriteValue(",")
		}
		cs.WriteValue(internal.ToString(v, nil))
	}
	cs.WriteValue(")")
	return cs, nil
}
