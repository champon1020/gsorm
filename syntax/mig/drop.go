package mig

import (
	"github.com/champon1020/mgorm/internal"
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

// BUild makes DROP DATABASE clause with syntax.StmtSet.
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

// DropIndex is DROP INDEX command.
type DropIndex struct {
	IdxName string
}

// Keyword returns clause keyword.
func (d *DropIndex) Keyword() string {
	return "DROP INDEX"
}

// Build makes DROP INDEX clause with syntax.StmtSet.
func (d *DropIndex) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.IdxName)
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
	Driver internal.SQLDriver
	Key    string
}

// Keyword returns clause keyword.
func (d *DropPrimary) Keyword() string {
	if d.Driver == internal.PSQL {
		return "DROP CONSTRAINT"
	}
	return "DROP PRIMARY KEY"
}

// Build makes DROP PRIMARY KEY | DROP CONSTRAINT clause.
func (d *DropPrimary) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	if d.Driver == internal.PSQL {
		ss.WriteValue(d.Key)
	}
	return ss, nil
}

// DropForeign is DROP FOREIGN KEY | DROP CONSTRAINT clause.
type DropForeign struct {
	Driver internal.SQLDriver
	Key    string
}

// Keyword returns clause keyword.
func (d *DropForeign) Keyword() string {
	if d.Driver == internal.PSQL {
		return "DROP CONSTRAINT"
	}
	return "DROP FOREIGN KEY"
}

// Build makes DROP FOREIGN KEY | DROP CONSTRAINT clause.
func (d *DropForeign) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Key)
	return ss, nil
}

// DropUnique is DROP UNIQUE | DROP CONSTRAINT clause.
type DropUnique struct {
	Driver internal.SQLDriver
	Key    string
}

// Keyword returns clause keyword.
func (d *DropUnique) Keyword() string {
	if d.Driver == internal.PSQL {
		return "DROP CONSTRAINT"
	}
	return "DROP INDEX"
}

// Build makes DROP INDEX KEY | DROP CONSTRAINT clause.
func (d *DropUnique) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Keyword())
	ss.WriteValue(d.Key)
	return ss, nil
}
