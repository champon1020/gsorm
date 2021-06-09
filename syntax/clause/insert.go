package clause

import (
	"fmt"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
)

// Insert is INSERT clause.
type Insert struct {
	Table   syntax.Table
	Columns []syntax.Column
}

// AddTable assigns the table to Insert.Table.
func (i *Insert) AddTable(table string) {
	i.Table = *syntax.NewTable(table)
}

// AddColumns appends the columns to Insert.Columns.
func (i *Insert) AddColumns(cols ...string) {
	for _, c := range cols {
		col := syntax.NewColumn(c)
		i.Columns = append(i.Columns, *col)
	}
}

// String returns function call as string.
func (i *Insert) String() string {
	s := fmt.Sprintf("%q", i.Table.Build())
	for _, c := range i.Columns {
		s += fmt.Sprintf(", %q", c.Build())
	}
	return fmt.Sprintf("Insert(%s)", s)
}

// Build creates the structure of INSERT clause that implements interfaces.ClauseSet.
func (i *Insert) Build() (interfaces.ClauseSet, error) {
	cs := &syntax.ClauseSet{}
	cs.WriteKeyword("INSERT INTO")
	cs.WriteValue(i.Table.Build())
	if len(i.Columns) > 0 {
		cs.WriteValue("(")
		for j, c := range i.Columns {
			if j != 0 {
				cs.WriteValue(",")
			}
			cs.WriteValue(c.Build())
		}
		cs.WriteValue(")")
	}
	return cs, nil
}
