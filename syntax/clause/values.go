package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Values is VALUES clause.
type Values struct {
	Values []interface{}
}

// Name returns clause keyword.
func (v *Values) Name() string {
	return "VALUES"
}

// AddValue appends values to Values.
func (v *Values) AddValue(val interface{}) {
	v.Values = append(v.Values, val)
}

// String returns function call with string.
func (v *Values) String() string {
	s := internal.SliceToString(v.Values)
	return fmt.Sprintf("%s(%s)", v.Name(), s)
}

// Build makes VALUES clause with sytnax.StmtSet.
func (v *Values) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(v.Name())
	ss.WriteValue("(")
	for i, v := range v.Values {
		if i != 0 {
			ss.WriteValue(",")
		}
		vStr, err := internal.ToString(v, true)
		if err != nil {
			return nil, err
		}
		ss.WriteValue(vStr)
	}
	ss.WriteValue(")")
	return ss, nil
}
