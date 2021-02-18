package mig

import "github.com/champon1020/mgorm/syntax"

// Check is CHECK clause.
type Check struct {
	Expr   string
	Values []interface{}
}

// Name returns clause keyword.
func (c *Check) Name() string {
	return "CHECK"
}

// Build makes CHECK clause with syntax.StmtSet.
func (c *Check) Build() (*syntax.StmtSet, error) {
	s, err := syntax.BuildForExpression(c.Expr, c.Values...)
	if err != nil {
		return nil, err
	}
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Name())
	ss.WriteValue("(")
	ss.WriteValue(s)
	ss.WriteValue(")")
	return ss, nil
}
