package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/mp/alter"
	"github.com/champon1020/mgorm/mp/create"
	"github.com/champon1020/mgorm/mp/drop"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/champon1020/mgorm/syntax/mig"
)

// New creates DB.
func New(dn, dsn string) (*DB, error) {
	db, err := sql.Open(dn, dsn)
	if err != nil {
		return nil, err
	}
	if dn == "psql" {
		return &DB{conn: db, driver: internal.PSQL}, nil
	}
	return &DB{conn: db, driver: internal.MySQL}, nil
}

// NewMock creates MockDB.
func NewMock() *MockDB {
	mock := new(MockDB)
	return mock
}

// Select calls SELECT command.
func Select(conn Conn, cols ...string) MgormSelect {
	s := new(SelectStmt)
	s.conn = conn
	s.cmd = clause.NewSelect(cols)
	return s
}

// Insert calls INSERT command.
func Insert(conn Conn, table string, cols ...string) MgormInsert {
	s := new(InsertStmt)
	s.conn = conn
	s.cmd = clause.NewInsert(table, cols)
	return s
}

// Update calls UPDATE command.
func Update(conn Conn, table string, cols ...string) MgormUpdate {
	s := new(UpdateStmt)
	s.conn = conn
	s.cmd = clause.NewUpdate(table, cols)
	return s
}

// Delete calls DELETE command.
func Delete(conn Conn) MgormDelete {
	s := new(DeleteStmt)
	s.conn = conn
	s.cmd = clause.NewDelete()
	return s
}

// Count calls COUNT function.
func Count(conn Conn, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.conn = conn
	return s
}

// Avg calls AVG function.
func Avg(conn Conn, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.conn = conn
	return s
}

// Sum calls SUM function.
func Sum(conn Conn, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.conn = conn
	return s
}

// Min calls MIN function.
func Min(conn Conn, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.conn = conn
	return s
}

// Max calls MAX function.
func Max(conn Conn, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.conn = conn
	return s
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn Conn, dbName string) create.DBMP {
	s := &CreateDBStmt{cmd: &mig.CreateDB{DBName: dbName}}
	s.conn = conn
	return s
}

// DropDB calls DROP DATABASE command.
func DropDB(conn Conn, dbName string) drop.DBMP {
	s := &DropDBStmt{cmd: &mig.DropDB{DBName: dbName}}
	s.conn = conn
	return s
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn Conn, table string) create.TableMP {
	s := &CreateTableStmt{cmd: &mig.CreateTable{Table: table}}
	s.conn = conn
	return s
}

// DropTable calls DROP TABLE command.
func DropTable(conn Conn, table string) drop.TableMP {
	s := &DropTableStmt{cmd: &mig.DropTable{Table: table}}
	s.conn = conn
	return s
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn Conn, table string) alter.TableMP {
	s := &AlterTableStmt{cmd: &mig.AlterTable{Table: table}}
	s.conn = conn
	return s
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn Conn, idx string) create.IndexMP {
	s := &CreateIndexStmt{cmd: &mig.CreateIndex{IdxName: idx}}
	s.conn = conn
	return s
}

// DropIndex calls DROP INDEX command.
func DropIndex(conn Conn, idx string) drop.IndexMP {
	s := &DropIndexStmt{cmd: &mig.DropIndex{IdxName: idx}}
	s.conn = conn
	return s
}
