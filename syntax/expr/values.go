package expr

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Values expression.
type Values struct {
	Columns []interface{}
}

// Name returns string of clause.
func (v *Values) Name() string {
	return "VALUES"
}

func (v *Values) addColumn(val interface{}) {
	v.Columns = append(v.Columns, val)
}

// String returns string of function call.
func (v *Values) String() string {
	s := internal.SliceToString(v.Columns)
	return fmt.Sprintf("%s(%s)", v.Name(), s)
}

// Build make values statement set.
func (v *Values) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteClause(v.Name())
	ss.WriteValue("(")
	for i, c := range v.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		cStr, err := internal.ToString(c)
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
