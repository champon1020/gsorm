package syntax

import "strconv"

// Offset expression.
type Offset struct {
	Num int
}

func (l *Offset) name() string {
	return "LIMIT"
}

// Build make limit statement set.
func (l *Offset) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(l.name())
	ss.WriteValue(strconv.Itoa(l.Num))
	return ss, nil
}

// NewOffset create new limit object.
func NewOffset(num int) *Offset {
	return &Offset{Num: num}
}
