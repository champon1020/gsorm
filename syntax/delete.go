package syntax

// Delete statement.
type Delete struct{}

func (d *Delete) query() string {
	return "DELETE"
}

// Build make delete statement set.
func (d *Delete) Build() *StmtSet {
	ss := new(StmtSet)
	ss.WriteClause(d.query())
	return ss
}

// NewDelete create new delete object.
func NewDelete() *Delete {
	return &Delete{}
}
