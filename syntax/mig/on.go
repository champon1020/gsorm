package mig

import "github.com/champon1020/mgorm/syntax"

// On is ON clause which is used with CREATE INDEX.
type On struct {
	Table   string
	Columns []string
}

// Name returns clause keyword.
func (o *On) Name() string {
	return "ON"
}

// Build makes ON clause with syntax.StmtSet.
func (o *On) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(o.Name())
	ss.WriteValue(o.Table)
	ss.WriteValue("(")
	for i, c := range o.Columns {
		if i > 0 {
			ss.WriteValue(",")
		}
		ss.WriteValue(c)
	}
	ss.WriteValue(")")
	return ss, nil
}
