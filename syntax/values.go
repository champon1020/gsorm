package syntax

import "github.com/champon1020/mgorm/internal"

// Values expression.
type Values struct {
	Columns []interface{}
}

func (v *Values) name() string {
	return "VALUES"
}

func (v *Values) addColumn(val interface{}) {
	v.Columns = append(v.Columns, val)
}

// Build make values statement set.
func (v *Values) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(v.name())
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
