package clause

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// Insert is INSERT clause.
type Insert struct {
	Table   syntax.Table
	Columns []syntax.Column
}

// Name returns clause keyword.
func (i *Insert) Name() string {
	return "INSERT INTO"
}

// AddTable appends table to Insert.
func (i *Insert) AddTable(table string) {
	i.Table = *syntax.NewTable(table)
}

// AddColumns appends columns to Insert.
func (i *Insert) AddColumns(cols ...string) {
	for _, c := range cols {
		col := syntax.NewColumn(c)
		i.Columns = append(i.Columns, *col)
	}
}

// String returns function call with string.
func (i *Insert) String() string {
	s := fmt.Sprintf("%q", i.Table.Build())
	for _, c := range i.Columns {
		s += fmt.Sprintf(", %q", c.Build())
	}
	return fmt.Sprintf("%s(%s)", i.Name(), s)
}

// Build makes INSERT clause with sytnax.StmtSet.
func (i *Insert) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(i.Name())
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
	return ss, nil
}
