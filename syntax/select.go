package syntax

import "fmt"

// Select statement.
type Select struct {
	Columns []Column
}

func (s *Select) query() string {
	return "SELECT"
}

func (s *Select) addColumn(col string) {
	c := NewColumn(col)
	s.Columns = append(s.Columns, *c)
}

// String returns string of function call.
func (s *Select) String() string {
	var str string
	for i, c := range s.Columns {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%q", c.Build())
	}
	return fmt.Sprintf("%s(%s)", s.query(), str)
}

// Build make select statement set.
func (s *Select) Build() *StmtSet {
	ss := &StmtSet{}
	ss.WriteClause(s.query())
	for i, c := range s.Columns {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	return ss
}

// NewSelect create new select object.
func NewSelect(cols []string) *Select {
	s := new(Select)
	for _, c := range cols {
		s.addColumn(c)
	}
	return s
}
