package mig

import "github.com/champon1020/mgorm/syntax"

// Constraint is CONSTRAINT clause.
type Constraint struct {
	Key string
}

// Name returns clause keyword.
func (c *Constraint) Name() string {
	return "CONSTRAINT"
}

// Build makes CONSTRAINT clasue with syntax.StmtSet.
func (c *Constraint) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Name())
	ss.WriteValue(c.Key)
	return ss, nil
}
