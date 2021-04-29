package mgorm

import (
	"fmt"

	"github.com/champon1020/mgorm/database"
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
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/statement/migration"
)

// Open opens the database connection.
func Open(dn, dsn string) (database.DB, error) {
	return database.NewDB(dn, dsn)
}

// OpenMock opens the mock database connection.
func OpenMock() database.MockDB {
	return database.NewMockDB()
}

// Select calls SELECT command.
func Select(conn database.Conn, cols ...string) ifselect.Stmt {
	return statement.NewSelectStmt(conn, cols...)
}

// Insert calls INSERT command.
func Insert(conn database.Conn, table string, cols ...string) ifinsert.Stmt {
	return statement.NewInsertStmt(conn, table, cols...)
}

// Update calls UPDATE command.
func Update(conn database.Conn, table string) ifupdate.Stmt {
	return statement.NewUpdateStmt(conn, table)
}

// Delete calls DELETE command.
func Delete(conn database.Conn) ifdelete.Stmt {
	return statement.NewDeleteStmt(conn)
}

// Count calls COUNT function.
func Count(conn database.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Avg calls AVG function.
func Avg(conn database.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Sum calls SUM function.
func Sum(conn database.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Min calls MIN function.
func Min(conn database.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Max calls MAX function.
func Max(conn database.Conn, col string, alias ...string) ifselect.Stmt {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn database.Conn, dbName string) ifcreatedb.DB {
	return migration.NewCreateDBStmt(conn, dbName)
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn database.Conn, idx string) ifcreateindex.Index {
	return migration.NewCreateIndexStmt(conn, idx)
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn database.Conn, table string) ifcreatetable.Table {
	return migration.NewCreateTableStmt(conn, table)
}

// DropDB calls DROP DATABASE command.
func DropDB(conn database.Conn, dbName string) ifdropdb.DB {
	return migration.NewDropDBStmt(conn, dbName)
}

// DropIndex calls DROP INDEX command.
func DropIndex(conn database.Conn, idx string) ifdropindex.Index {
	return migration.NewDropIndexStmt(conn, idx)
}

// DropTable calls DROP TABLE command.
func DropTable(conn database.Conn, table string) ifdroptable.Table {
	return migration.NewDropTableStmt(conn, table)
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn database.Conn, table string) ifaltertable.Table {
	return migration.NewAlterTableStmt(conn, table)
}
