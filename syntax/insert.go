package syntax

// Insert statement.
type Insert struct {
	Table   Table
	Columns []Column
}

func (i *Insert) query() string {
	return "INSERT INTO"
}

func (i *Insert) addTable(table string) {
	i.Table = *NewTable(table)
}

func (i *Insert) addColumn(col string) {
	column := NewColumn(col)
	i.Columns = append(i.Columns, *column)
}

// Build make insert statement set.
func (i *Insert) Build() *StmtSet {
	ss := new(StmtSet)
	ss.WriteClause(i.query())
	ss.WriteValue(i.Table.Build())
	ss.WriteValue("(")
	for j, c := range i.Columns {
		if j != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c.Build())
	}
	ss.WriteValue(")")
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
