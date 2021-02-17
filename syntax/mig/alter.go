package mig

import "github.com/champon1020/mgorm/syntax"

// AlterTable is ALTER TABLE clause.
type AlterTable struct {
	Table string
}

func (a *AlterTable) Name() string {
	return "ALTER TABLE"
}

func (a *AlterTable) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(a.Name())
	ss.WriteValue(a.Table)
	return ss
}
