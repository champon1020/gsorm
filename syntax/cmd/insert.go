package cmd

import (
	"fmt"

	"github.com/champon1020/mgorm/syntax"
)

// Insert statement.
type Insert struct {
	Table   syntax.Table
	Columns []syntax.Column
}

// Query returns string of clause.
func (i *Insert) Query() string {
	return "INSERT INTO"
}

func (i *Insert) addTable(table string) {
	i.Table = *syntax.NewTable(table)
}

func (i *Insert) addColumn(col string) {
	column := syntax.NewColumn(col)
	i.Columns = append(i.Columns, *column)
}

// String returns string of function call.
func (i *Insert) String() string {
	s := fmt.Sprintf("%q", i.Table.Build())
	for _, c := range i.Columns {
		s += fmt.Sprintf(", %q", c.Build())
	}
	return fmt.Sprintf("%s(%s)", i.Query(), s)
}

// Build make insert statement set.
func (i *Insert) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteClause(i.Query())
	ss.WriteValue(i.Table.Build())
	if len(i.Columns) > 0 {
		ss.WriteValue("(")
		for j, c := range i.Columns {
			if j != 0 {
				ss.WriteValue(",")
			}
			ss.WriteValue(c.Build())
		}
		ss.WriteValue(")")
	}
	return ss
}

// NewInsert create new insert object.
func NewInsert(table string, cols []string) *Insert {
	i := new(Insert)
	i.addTable(table)
	for _, c := range cols {
		i.addColumn(c)
	}
	return i
}
