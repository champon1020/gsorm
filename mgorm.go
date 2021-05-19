package mgorm

import (
	"fmt"

	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/interfaces/ialtertable"
	"github.com/champon1020/mgorm/interfaces/icreatedb"
	"github.com/champon1020/mgorm/interfaces/icreateindex"
	"github.com/champon1020/mgorm/interfaces/icreatetable"
	"github.com/champon1020/mgorm/interfaces/idelete"
	"github.com/champon1020/mgorm/interfaces/idropdb"
	"github.com/champon1020/mgorm/interfaces/idroptable"
	"github.com/champon1020/mgorm/interfaces/iinsert"
	"github.com/champon1020/mgorm/interfaces/iraw"
	"github.com/champon1020/mgorm/interfaces/iselect"
	"github.com/champon1020/mgorm/interfaces/iupdate"
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

// RawStmt calls raw string statement.
func RawStmt(conn domain.Conn, rs string, v ...interface{}) iraw.Stmt {
	return statement.NewRawStmt(conn, rs, v...)
}

// Select calls SELECT command.
func Select(conn domain.Conn, cols ...string) iselect.Stmt {
	return statement.NewSelectStmt(conn, cols...)
}

// Insert calls INSERT command.
func Insert(conn domain.Conn, table string, cols ...string) iinsert.Stmt {
	return statement.NewInsertStmt(conn, table, cols...)
}

// Update calls UPDATE command.
func Update(conn domain.Conn, table string) iupdate.Stmt {
	return statement.NewUpdateStmt(conn, table)
}

// Delete calls DELETE command.
func Delete(conn domain.Conn) idelete.Stmt {
	return statement.NewDeleteStmt(conn)
}

// Count calls COUNT function.
func Count(conn domain.Conn, col string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Avg calls AVG function.
func Avg(conn domain.Conn, col string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Sum calls SUM function.
func Sum(conn domain.Conn, col string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Min calls MIN function.
func Min(conn domain.Conn, col string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Max calls MAX function.
func Max(conn domain.Conn, col string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn domain.Conn, dbName string) icreatedb.DB {
	return migration.NewCreateDBStmt(conn, dbName)
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn domain.Conn, idx string) icreateindex.Index {
	return migration.NewCreateIndexStmt(conn, idx)
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn domain.Conn, table string) icreatetable.Table {
	return migration.NewCreateTableStmt(conn, table)
}

// DropDB calls DROP DATABASE command.
func DropDB(conn domain.Conn, dbName string) idropdb.DB {
	return migration.NewDropDBStmt(conn, dbName)
}

// DropTable calls DROP TABLE command.
func DropTable(conn domain.Conn, table string) idroptable.Table {
	return migration.NewDropTableStmt(conn, table)
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn domain.Conn, table string) ialtertable.Table {
	return migration.NewAlterTableStmt(conn, table)
}
