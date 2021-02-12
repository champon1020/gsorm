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
	var (
		name  string
		alias string
	)

	strs := strings.Split(table, " AS ")
	name = strs[0]
	if len(strs) == 2 {
		alias = strs[1]
	}

	return &Table{Name: name, Alias: alias}
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
	var (
		name  string
		alias string
	)

	strs := strings.Split(column, " AS ")
	name = strs[0]
	if len(strs) == 2 {
		alias = strs[1]
	}

	return &Column{Name: name, Alias: alias}
}

// Eq is equal expression.
type Eq struct {
	LHS string
	RHS interface{}
}

// Build makes equal expression with string.
func (e *Eq) Build() (string, error) {
	s := e.LHS
	s += " = "
	rhs, err := internal.ToString(e.RHS, true)
	s += rhs
	return s, err
}

// NewEq creates new Eq instance.
func NewEq(lhs string, rhs interface{}) *Eq {
	return &Eq{LHS: lhs, RHS: rhs}
}
