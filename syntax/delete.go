package syntax

import "fmt"

// Delete statement.
type Delete struct{}

func (d *Delete) query() string {
	return "DELETE"
}

// String returns string of function call.
func (d *Delete) String() string {
	return fmt.Sprintf("%s()", d.query())
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
