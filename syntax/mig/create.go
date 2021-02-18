package mig

import "github.com/champon1020/mgorm/syntax"

// CreateDB is CREATE DATABASE clause.
type CreateDB struct {
	DBName string
}

// Name returns clause keyword.
func (c *CreateDB) Query() string {
	return "CREATE DATABASE"
}

// Build makes CREATE DATABASE clause with syntax.StmtSet.
func (c *CreateDB) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Query())
	ss.WriteValue(c.DBName)
	return ss
}

// CreateTable is CREATE TABLE clause.
type CreateTable struct {
	Table string
}

// Name returns clause keyword.
func (c *CreateTable) Query() string {
	return "CREATE TABLE"
}

// Build makes CREATE TABLE clause with syntax.StmtSet.
func (c *CreateTable) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Query())
	ss.WriteValue(c.Table)
	return ss
}

// CreateIndex is CREATE INDEX clause.
type CreateIndex struct {
	IdxName string
}

// Query returns clause keyword.
func (c *CreateIndex) Query() string {
	return "CREATE INDEX"
}

// Build makes CREATE INDEX clause with syntax.StmtSet.
func (c *CreateIndex) Build() *syntax.StmtSet {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Query())
	ss.WriteValue(c.IdxName)
	return ss
}
