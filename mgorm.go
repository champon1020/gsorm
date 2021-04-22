package mgorm

import (
	"database/sql"
	"fmt"

	ifaltertable "github.com/champon1020/mgorm/interfaces/altertable"
	ifcreatedb "github.com/champon1020/mgorm/interfaces/createdb"
	ifcreateindex "github.com/champon1020/mgorm/interfaces/createindex"
	ifcreatetable "github.com/champon1020/mgorm/interfaces/createtable"
	ifdelete "github.com/champon1020/mgorm/interfaces/delete"
	ifdropdb "github.com/champon1020/mgorm/interfaces/dropdb"
	ifdropindex "github.com/champon1020/mgorm/interfaces/dropindex"
	ifdroptable "github.com/champon1020/mgorm/interfaces/droptable"
	ifinsert "github.com/champon1020/mgorm/interfaces/insert"
	ifselect "github.com/champon1020/mgorm/interfaces/select"
	ifupdate "github.com/champon1020/mgorm/interfaces/update"
	"github.com/champon1020/mgorm/internal"
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
func Select(conn Conn, cols ...string) ifselect.Stmt {
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
func Insert(conn Conn, table string, cols ...string) ifinsert.Stmt {
	i := new(clause.Insert)
	i.AddTable(table)
	i.AddColumns(cols...)
	s := &InsertStmt{cmd: i}
	s.conn = conn
	return s
}

// Update calls UPDATE command.
func Update(conn Conn, table string) ifupdate.Stmt {
	u := new(clause.Update)
	u.AddTable(table)
	s := &UpdateStmt{cmd: u}
	s.conn = conn
	return s
}

// Delete calls DELETE command.
func Delete(conn Conn) ifdelete.Stmt {
	s := &DeleteStmt{cmd: &clause.Delete{}}
	s.conn = conn
	return s
}

// Count calls COUNT function.
func Count(conn Conn, col string, alias ...string) ifselect.Stmt {
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
func Avg(conn Conn, col string, alias ...string) ifselect.Stmt {
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
func Sum(conn Conn, col string, alias ...string) ifselect.Stmt {
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
func Min(conn Conn, col string, alias ...string) ifselect.Stmt {
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
func Max(conn Conn, col string, alias ...string) ifselect.Stmt {
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
func CreateDB(conn Conn, dbName string) ifcreatedb.DB {
	s := &CreateDBStmt{cmd: &mig.CreateDB{DBName: dbName}}
	s.conn = conn
	return s
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn Conn, idx string) ifcreateindex.Index {
	s := &CreateIndexStmt{cmd: &mig.CreateIndex{IdxName: idx}}
	s.conn = conn
	return s
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn Conn, table string) ifcreatetable.Table {
	s := &CreateTableStmt{cmd: &mig.CreateTable{Table: table}}
	s.conn = conn
	return s
}

// DropDB calls DROP DATABASE command.
func DropDB(conn Conn, dbName string) ifdropdb.DB {
	s := &DropDBStmt{cmd: &mig.DropDB{DBName: dbName}}
	s.conn = conn
	return s
}

// DropIndex calls DROP INDEX command.
func DropIndex(conn Conn, idx string) ifdropindex.Index {
	s := &DropIndexStmt{cmd: &mig.DropIndex{IdxName: idx}}
	s.conn = conn
	return s
}

// DropTable calls DROP TABLE command.
func DropTable(conn Conn, table string) ifdroptable.Table {
	s := &DropTableStmt{cmd: &mig.DropTable{Table: table}}
	s.conn = conn
	return s
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn Conn, table string) ifaltertable.Table {
	s := &AlterTableStmt{cmd: &mig.AlterTable{Table: table}}
	s.conn = conn
	return s
}
