package syntax

// Update statement.
type Update struct {
	Table Table
}

func (u *Update) query() string {
	return "UPDATE"
}

func (u *Update) addTable(table string) {
	u.Table = *NewTable(table)
}

// Build make update statement set.
func (u *Update) Build() *StmtSet {
	ss := new(StmtSet)
	ss.WriteClause(u.query())
	ss.WriteValue(u.Table.Build())
	return ss
}

// NewUpdate create new update object.
func NewUpdate(table string) *Update {
	u := new(Update)
	u.addTable(table)
	return u
}
