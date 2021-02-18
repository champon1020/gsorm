package mig

import "github.com/champon1020/mgorm/syntax"

// Charset is CHARSET clause.
type Charset struct {
	Format string
}

// Name returns clause keyword.
func (c *Charset) Name() string {
	return "CHARSET"
}

// Build makes CHARSET clause with syntax.StmtSet.
func (c *Charset) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Name())
	ss.WriteValue("=")
	ss.WriteValue(c.Format)
	return ss, nil
}
