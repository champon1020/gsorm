package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Values is VALUES clause.
type Values struct {
	Columns []interface{}
}

// Name returns clause keyword.
func (v *Values) Name() string {
	return "VALUES"
}

// addColumn appends values to Values.
func (v *Values) addColumn(val interface{}) {
	v.Columns = append(v.Columns, val)
}

// String returns function call with string.
func (v *Values) String() string {
	s := internal.SliceToString(v.Columns)
	return fmt.Sprintf("%s(%s)", v.Name(), s)
}

// Build makes VALUES clause with sytnax.StmtSet.
func (v *Values) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(v.Name())
	ss.WriteValue("(")
	for i, c := range v.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		cStr, err := internal.ToString(c, true)
		if err != nil {
			return nil, err
		}
		ss.WriteValue(cStr)
	}
	ss.WriteValue(")")
	return ss, nil
}

// NewValues create new values object.
func NewValues(cols []interface{}) *Values {
	v := new(Values)
	for _, c := range cols {
		v.addColumn(c)
	}
	return v
}
