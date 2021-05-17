package mig

import (
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

// Build makes DROP DATABASE clause with syntax.StmtSet.
func (d *DropDB) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.DBName)
	return ss, nil
}

// DropTable is DROP TABLE clause.
type DropTable struct {
	Table string
}

// Keyword returns clause keyword.
func (d *DropTable) Keyword() string {
	return "DROP TABLE"
}

// Build makes DROP TABLE clause with syntax.StmtSet.
func (d *DropTable) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Table)
	return ss, nil
}

// DropColumn is DROP COLUMN clause.
type DropColumn struct {
	Column string
}

// Keyword returns clause keyword.
func (d *DropColumn) Keyword() string {
	return "DROP COLUMN"
}

// Build makes DROP COLUMN clause with syntax.StmtSet.
func (d *DropColumn) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Column)
	return ss, nil
}

// DropPrimary is DROP PRIMARY KEY | DROP CONSTRAINT clause.
type DropPrimary struct {
	Driver domain.SQLDriver
	Key    string
}
