package syntax

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
