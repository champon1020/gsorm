package mgorm

import (
	"fmt"

	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/domain"
	ifaltertable "github.com/champon1020/mgorm/interfaces/altertable"
	ifcreatedb "github.com/champon1020/mgorm/interfaces/createdb"
	ifcreateindex "github.com/champon1020/mgorm/interfaces/createindex"
	ifcreatetable "github.com/champon1020/mgorm/interfaces/createtable"
	ifdelete "github.com/champon1020/mgorm/interfaces/delete"
	ifdropdb "github.com/champon1020/mgorm/interfaces/dropdb"
	ifdroptable "github.com/champon1020/mgorm/interfaces/droptable"
	ifinsert "github.com/champon1020/mgorm/interfaces/insert"
	ifselect "github.com/champon1020/mgorm/interfaces/select"
	ifupdate "github.com/champon1020/mgorm/interfaces/update"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/statement/migration"
)

// Open opens the database connection.
func Open(driver, dsn string) (domain.DB, error) {
	return database.NewDB(driver, dsn)
}

// OpenMock opens the mock database connection.
func OpenMock(driver string) domain.MockDB {
	return database.NewMockDB(driver)
}

// Select calls SELECT command.
func Select(conn domain.Conn, cols ...string) ifselect.Stmt {
	return statement.NewSelectStmt(conn, cols...)
}

// Insert calls INSERT command.
func Insert(conn domain.Conn, table string, cols ...string) ifinsert.Stmt {
	return statement.NewInsertStmt(conn, table, cols...)
}

// Update calls UPDATE command.
func Update(conn domain.Conn, table string) ifupdate.Stmt {
	return statement.NewUpdateStmt(conn, table)
}

// Delete calls DELETE command.
func Delete(conn domain.Conn) ifdelete.Stmt {
	return statement.NewDeleteStmt(conn)
}

// Count calls COUNT function.
func Count(conn domain.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Avg calls AVG function.
func Avg(conn domain.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Sum calls SUM function.
func Sum(conn domain.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Min calls MIN function.
func Min(conn domain.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Max calls MAX function.
func Max(conn domain.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn domain.Conn, dbName string) ifcreatedb.DB {
	return migration.NewCreateDBStmt(conn, dbName)
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn domain.Conn, idx string) ifcreateindex.Index {
	return migration.NewCreateIndexStmt(conn, idx)
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn domain.Conn, table string) ifcreatetable.Table {
	return migration.NewCreateTableStmt(conn, table)
}

// DropDB calls DROP DATABASE command.
func DropDB(conn domain.Conn, dbName string) ifdropdb.DB {
	return migration.NewDropDBStmt(conn, dbName)
}

// DropTable calls DROP TABLE command.
func DropTable(conn domain.Conn, table string) ifdroptable.Table {
	return migration.NewDropTableStmt(conn, table)
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn domain.Conn, table string) ifaltertable.Table {
	return migration.NewAlterTableStmt(conn, table)
}
