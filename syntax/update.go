package syntax

import "fmt"

// Update statement.
type Update struct {
	Table   Table
	Columns []string
}

func (u *Update) query() string {
	return "UPDATE"
}

func (u *Update) addTable(table string) {
	u.Table = *NewTable(table)
}

func (u *Update) addColumns(cols []string) {
	u.Columns = cols
}

// String returns string of function call.
func (u *Update) String() string {
	s := fmt.Sprintf("%q", u.Table.Build())
	for _, c := range u.Columns {
		s += fmt.Sprintf(", %q", c)
	}
	return fmt.Sprintf("%s(%s)", u.query(), s)
}

// Build make update statement set.
func (u *Update) Build() *StmtSet {
	ss := new(StmtSet)
	ss.WriteClause(u.query())
	ss.WriteValue(u.Table.Build())
	return ss
}

// NewUpdate create new update object.
func NewUpdate(table string, cols []string) *Update {
	u := new(Update)
	u.addTable(table)
	u.addColumns(cols)
	return u
}
