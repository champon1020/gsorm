package syntax

import (
	"strings"

	"github.com/champon1020/mgorm/internal"
)

// Table expression.
type Table struct {
	Name  string
	Alias string
}

// Build make table expression as string.
func (t *Table) Build() string {
	s := t.Name
	if len(t.Alias) > 0 {
		s += " AS "
	}
	s += t.Alias
	return s
}

// NewTable generate the new table object.
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

// Column expression.
type Column struct {
	Name  string
	Alias string
}

// Build make column expression as string.
func (c *Column) Build() string {
	s := c.Name
	if len(c.Alias) > 0 {
		s += " AS "
	}
	s += c.Alias
	return s
}

// NewColumn generate the new column object.
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

// Eq expression.
type Eq struct {
	LHS string
	RHS interface{}
}

// Build make equal expression as string.
func (e *Eq) Build() (string, error) {
	s := e.LHS
	s += " = "
	rhs, err := internal.ToString(e.RHS)
	s += rhs
	return s, err
}

// NewEq generate the new eq object.
func NewEq(lhs string, rhs interface{}) *Eq {
	return &Eq{LHS: lhs, RHS: rhs}
}
