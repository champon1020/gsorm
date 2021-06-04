package gsorm

import (
	"fmt"

	"github.com/champon1020/gsorm/database"
	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/ialtertable"
	"github.com/champon1020/gsorm/interfaces/icreatedb"
	"github.com/champon1020/gsorm/interfaces/icreateindex"
	"github.com/champon1020/gsorm/interfaces/icreatetable"
	"github.com/champon1020/gsorm/interfaces/idelete"
	"github.com/champon1020/gsorm/interfaces/idropdb"
	"github.com/champon1020/gsorm/interfaces/idroptable"
	"github.com/champon1020/gsorm/interfaces/iinsert"
	"github.com/champon1020/gsorm/interfaces/iraw"
	"github.com/champon1020/gsorm/interfaces/iselect"
	"github.com/champon1020/gsorm/interfaces/iupdate"
	"github.com/champon1020/gsorm/statement"
	"github.com/champon1020/gsorm/statement/migration"
)

type (
	// DB interface.
	DB = domain.DB

	// Tx interface.
	Tx = domain.Tx

	// MockDB interface.
	MockDB = domain.MockDB

	// MockTx interface.
	MockTx = domain.MockTx
)

// Open opens the database connection.
func Open(driver, dsn string) (DB, error) {
	return database.NewDB(driver, dsn)
}

// OpenMock opens the mock database connection.
func OpenMock() MockDB {
	return database.NewMockDB()
}

// RawStmt calls raw string statement.
func RawStmt(conn domain.Conn, raw string, values ...interface{}) iraw.Stmt {
	return statement.NewRawStmt(conn, raw, values...)
}

// Select calls SELECT command.
func Select(conn domain.Conn, columns ...string) iselect.Stmt {
	return statement.NewSelectStmt(conn, columns...)
}

// Insert calls INSERT command.
func Insert(conn domain.Conn, table string, columns ...string) iinsert.Stmt {
	return statement.NewInsertStmt(conn, table, columns...)
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
func Count(conn domain.Conn, column string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("COUNT(%s)", column)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Avg calls AVG function.
func Avg(conn domain.Conn, column string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("AVG(%s)", column)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Sum calls SUM function.
func Sum(conn domain.Conn, column string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("SUM(%s)", column)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Min calls MIN function.
func Min(conn domain.Conn, column string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("MIN(%s)", column)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// Max calls MAX function.
func Max(conn domain.Conn, column string, alias ...string) iselect.Stmt {
	c := fmt.Sprintf("MAX(%s)", column)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	return statement.NewSelectStmt(conn, c)
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn domain.Conn, dbName string) icreatedb.Stmt {
	return migration.NewCreateDBStmt(conn, dbName)
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn domain.Conn, idx string) icreateindex.Stmt {
	return migration.NewCreateIndexStmt(conn, idx)
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn domain.Conn, table string) icreatetable.Stmt {
	return migration.NewCreateTableStmt(conn, table)
}

// DropDB calls DROP DATABASE command.
func DropDB(conn domain.Conn, dbName string) idropdb.Stmt {
	return migration.NewDropDBStmt(conn, dbName)
}

// DropTable calls DROP TABLE command.
func DropTable(conn domain.Conn, table string) idroptable.Stmt {
	return migration.NewDropTableStmt(conn, table)
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn domain.Conn, table string) ialtertable.Stmt {
	return migration.NewAlterTableStmt(conn, table)
}
