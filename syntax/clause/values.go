package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
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
	s := internal.ToString(v.Values, nil)
	return fmt.Sprintf("%s(%s)", v.Name(), s)
}

// Build makes VALUES clause with sytnax.StmtSet.
func (v *Values) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(v.Name())
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
