package mig

import "github.com/champon1020/mgorm/syntax"

// ForeignKey is FOREIGN KEY clasue.
type ForeignKey struct {
	Column string
}

// Name returns clause keyword.
func (f *ForeignKey) Name() string {
	return "FOREIGN KEY"
}

// Build makes FOREIGN KEY clasue with syntax.StmtSet.
func (f *ForeignKey) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(f.Name())
	ss.WriteValue("(")
	ss.WriteValue(f.Column)
	ss.WriteValue(")")
	return ss, nil
}
