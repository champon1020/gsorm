package syntax

// From expression.
type From struct {
	Tables []Table
}

func (f *From) name() string {
	return "FROM"
}

func (f *From) addTable(col string) {
	c := NewTable(col)
	f.Tables = append(f.Tables, *c)
}

// Build make from statement set.
func (f *From) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(f.name())
	for i, t := range f.Tables {
		if i != 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(t.Build())
	}
	return ss, nil
}

// NewFrom make new from object.
func NewFrom(tables []string) *From {
	f := new(From)
	for _, t := range tables {
		f.addTable(t)
	}
	return f
}
