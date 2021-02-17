package mig

import "github.com/champon1020/mgorm/syntax"

// FK is FOREIGN KEY clasue.
type FK struct {
	Column string
}

// Name returns clause keyword.
func (f *FK) Name() string {
	return "FOREIGN KEY"
}

// Build makes FOREIGN KEY clasue with syntax.StmtSet.
func (f *FK) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(f.Name())
	ss.WriteValue("(")
	ss.WriteValue(f.Column)
	ss.WriteValue(")")
	return ss, nil
}
