package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/internal"
	prAlter "github.com/champon1020/mgorm/provider/alter"
	prCreate "github.com/champon1020/mgorm/provider/create"
	prDelete "github.com/champon1020/mgorm/provider/delete"
	prDrop "github.com/champon1020/mgorm/provider/drop"
	prInsert "github.com/champon1020/mgorm/provider/insert"
	prSelect "github.com/champon1020/mgorm/provider/select"
	prUpdate "github.com/champon1020/mgorm/provider/update"
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
func Select(conn Conn, cols ...string) prSelect.StmtMP {
	sel := new(clause.Select)
	if len(cols) == 0 {
		sel.AddColumns("*")
	} else {
		sel.AddColumns(cols...)
	}
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// Insert calls INSERT command.
func Insert(conn Conn, table string, cols ...string) prInsert.StmtMP {
	i := new(clause.Insert)
	i.AddTable(table)
	i.AddColumns(cols...)
	s := &InsertStmt{cmd: i}
	s.conn = conn
	return s
}

// Update calls UPDATE command.
func Update(conn Conn, table string, cols ...string) prUpdate.StmtMP {
	u := new(clause.Update)
	u.AddTable(table)
	u.AddColumns(cols)
	s := &UpdateStmt{cmd: u}
	s.conn = conn
	return s
}

// Delete calls DELETE command.
func Delete(conn Conn) prDelete.StmtMP {
	s := &DeleteStmt{cmd: &clause.Delete{}}
	s.conn = conn
	return s
}

// Count calls COUNT function.
func Count(conn Conn, col string, alias ...string) prSelect.StmtMP {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	sel := new(clause.Select)
	sel.AddColumns(c)
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// Avg calls AVG function.
func Avg(conn Conn, col string, alias ...string) prSelect.StmtMP {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	sel := new(clause.Select)
	sel.AddColumns(c)
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// Sum calls SUM function.
func Sum(conn Conn, col string, alias ...string) prSelect.StmtMP {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	sel := new(clause.Select)
	sel.AddColumns(c)
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// Min calls MIN function.
func Min(conn Conn, col string, alias ...string) prSelect.StmtMP {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	sel := new(clause.Select)
	sel.AddColumns(c)
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// Max calls MAX function.
func Max(conn Conn, col string, alias ...string) prSelect.StmtMP {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	sel := new(clause.Select)
	sel.AddColumns(c)
	s := &SelectStmt{cmd: sel}
	s.conn = conn
	return s
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn Conn, dbName string) prCreate.DBMP {
	s := &CreateDBStmt{cmd: &mig.CreateDB{DBName: dbName}}
	s.conn = conn
	return s
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn Conn, idx string) prCreate.IndexMP {
	s := &CreateIndexStmt{cmd: &mig.CreateIndex{IdxName: idx}}
	s.conn = conn
	return s
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn Conn, table string) prCreate.TableMP {
	s := &CreateTableStmt{cmd: &mig.CreateTable{Table: table}}
	s.conn = conn
	return s
}

// DropDB calls DROP DATABASE command.
func DropDB(conn Conn, dbName string) prDrop.DBMP {
	s := &DropDBStmt{cmd: &mig.DropDB{DBName: dbName}}
	s.conn = conn
	return s
}

// DropIndex calls DROP INDEX command.
func DropIndex(conn Conn, idx string) prDrop.IndexMP {
	s := &DropIndexStmt{cmd: &mig.DropIndex{IdxName: idx}}
	s.conn = conn
	return s
}

// DropTable calls DROP TABLE command.
func DropTable(conn Conn, table string) prDrop.TableMP {
	s := &DropTableStmt{cmd: &mig.DropTable{Table: table}}
	s.conn = conn
	return s
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn Conn, table string) prAlter.TableMP {
	s := &AlterTableStmt{cmd: &mig.AlterTable{Table: table}}
	s.conn = conn
	return s
}
