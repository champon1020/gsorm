package mig

import "github.com/champon1020/mgorm/syntax"

// DropDB is DROP DATABASE clause.
type DropDB struct {
	DBName string
}

// Query returns clause keyword.
func (d *DropDB) Query() string {
	return "DROP DATABASE"
}

// BUild makes DROP DATABASE clause with syntax.StmtSet.
func (d *DropDB) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Query())
	ss.WriteValue(d.DBName)
	return ss
}

// DropTable is DROP TABLE clause.
type DropTable struct {
	Table string
}

// Query returns clause keyword.
func (d *DropTable) Query() string {
	return "DROP TABLE"
}

// Build makes DROP TABLE clause with syntax.StmtSet.
func (d *DropTable) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Query())
	ss.WriteValue(d.Table)
	return ss
}

// Drop is DROP clause.
type Drop struct {
	Column string
}

// Name returns clause keyword.
func (d *Drop) Name() string {
	return "DROP"
}

// Build makes DROP clause with syntax.StmtSet.
func (d *Drop) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Name())
	ss.WriteValue(d.Column)
	return ss, nil
}

// Drop is DROP CONSTRAINT clause.
type DropConstraint struct {
	Key string
}

// Name returns clause keyword.
func (d *DropConstraint) Name() string {
	return "DROP CONSTRAINT"
}

// Build makes DROP CONSTRAINT clause with syntax.StmtSet.
func (d *DropConstraint) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Name())
	ss.WriteValue(d.Key)
	return ss, nil
}

// DropIndex is DROP INDEX
type DropIndex struct {
	IdxName string
}

// Query returns caluse keyword.
func (d *DropIndex) Name() string {
	return "DROP INDEX"
}

// Build makes DROP INDEX clause with syntax.StmtSet.
func (d *DropIndex) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(d.Name())
	ss.WriteValue(d.IdxName)
	return ss, nil
}
