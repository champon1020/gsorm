package syntax

import (
	"strings"

	"github.com/champon1020/mgorm/internal"
)

// Table is table term.
type Table struct {
	Name  string
	Alias string
}

// Build makes table term with string.
func (t *Table) Build() string {
	s := t.Name
	if len(t.Alias) > 0 {
		s += " AS "
	}
	s += t.Alias
	return s
}

// NewTable creates new Table instance.
func NewTable(table string) *Table {
	t := new(Table)
	t.Name = table
	if strs := strings.Split(table, " AS "); len(strs) == 2 {
		t.Name = strs[0]
		t.Alias = strs[1]
	}
	if strs := strings.Split(table, " as "); len(strs) == 2 {
		t.Name = strs[0]
		t.Alias = strs[1]
	}
	return t
}

// Column is column term.
type Column struct {
	Name  string
	Alias string
}

// Build makes column term with string.
func (c *Column) Build() string {
	s := c.Name
	if len(c.Alias) > 0 {
		s += " AS "
	}
	s += c.Alias
	return s
}

// NewColumn creates new Column instance.
func NewColumn(column string) *Column {
	c := new(Column)
	c.Name = column
	if strs := strings.Split(column, " AS "); len(strs) == 2 {
		c.Name = strs[0]
		c.Alias = strs[1]
	}
	if strs := strings.Split(column, " as "); len(strs) == 2 {
		c.Name = strs[0]
		c.Alias = strs[1]
	}
	return c
}

// Eq is equal expression.
type Eq struct {
	LHS string
	RHS interface{}
}

// Build makes equal expression with string.
func (e *Eq) Build() string {
	s := e.LHS
	s += " = "
	s += internal.ToString(e.RHS, true)
	return s
}

// NewEq creates new Eq instance.
func NewEq(lhs string, rhs interface{}) *Eq {
	return &Eq{LHS: lhs, RHS: rhs}
}
