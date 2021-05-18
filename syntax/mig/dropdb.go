package mig

import (
	"fmt"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/syntax"
)

// DropDB is DROP DATABASE clause.
type DropDB struct {
	DBName string
}

// Keyword returns clause keyword.
func (d *DropDB) Keyword() string {
	return "DROP DATABASE"
}

func (d *DropDB) String() string {
	return fmt.Sprintf("%s(%s)", d.Keyword(), d.DBName)
}

// Build makes DROP DATABASE clause with syntax.StmtSet.
func (d *DropDB) Build() (domain.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.DBName)
	return ss, nil
}
