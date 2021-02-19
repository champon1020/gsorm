package mig

import "github.com/champon1020/mgorm/syntax"

// CreateDB is CREATE DATABASE clause.
type CreateDB struct {
	DBName string
}

// Keyword returns clause keyword.
func (c *CreateDB) Keyword() string {
	return "CREATE DATABASE"
}

// Build makes CREATE DATABASE clause with syntax.StmtSet.
func (c *CreateDB) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.DBName)
	return ss, nil
}

// CreateTable is CREATE TABLE clause.
type CreateTable struct {
	Table string
}

// Keyword returns clause keyword.
func (c *CreateTable) Keyword() string {
	return "CREATE TABLE"
}

// Build makes CREATE TABLE clause with syntax.StmtSet.
func (c *CreateTable) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.Table)
	return ss, nil
}

// CreateIndex is CREATE INDEX clause.
type CreateIndex struct {
	IdxName string
}

// Keyword returns clause keyword.
func (c *CreateIndex) Keyword() string {
	return "CREATE INDEX"
}

// Build makes CREATE INDEX clause with syntax.StmtSet.
func (c *CreateIndex) Build() (*syntax.StmtSet, error) {
	ss := new(syntax.StmtSet)
	ss.WriteKeyword(c.Keyword())
	ss.WriteValue(c.IdxName)
	return ss, nil
}
