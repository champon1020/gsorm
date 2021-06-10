package gsorm

import (
	"database/sql"
	"fmt"

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
)

// Open opens the database connection.
func Open(driver, dsn string) (DB, error) {
	d, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &db{conn: d}, nil
}

// OpenMock opens the mock database connection.
func OpenMock() MockDB {
	return &mockDB{}
}

// RawStmt calls raw string statement.
func RawStmt(conn conn, raw string, values ...interface{}) iraw.Stmt {
	return newRawStmt(conn, raw, values...)
}

// Select calls SELECT command.
func Select(conn conn, columns ...string) iselect.Stmt {
	return newSelectStmt(conn, columns...)
}

// Insert calls INSERT command.
func Insert(conn conn, table string, columns ...string) iinsert.Stmt {
	return newInsertStmt(conn, table, columns...)
}

// Update calls UPDATE command.
func Update(conn conn, table string) iupdate.Stmt {
	return newUpdateStmt(conn, table)
}

// Delete calls DELETE command.
func Delete(conn conn) idelete.Stmt {
	return newDeleteStmt(conn)
}

// Count calls COUNT function.
func Count(conn conn, columns ...string) iselect.Stmt {
	if len(columns) > 0 {
		var cols string
		for i, c := range columns {
			if i > 0 {
				cols += ", "
			}
			cols += fmt.Sprintf("COUNT(%s)", c)
		}
		return newSelectStmt(conn, cols)
	}
	return newSelectStmt(conn)
}

// Sum calls SUM function.
func Sum(conn conn, columns ...string) iselect.Stmt {
	if len(columns) > 0 {
		var cols string
		for i, c := range columns {
			if i > 0 {
				cols += ", "
			}
			cols += fmt.Sprintf("SUM(%s)", c)
		}
		return newSelectStmt(conn, cols)
	}
	return newSelectStmt(conn)
}

// Avg calls AVG function.
func Avg(conn conn, columns ...string) iselect.Stmt {
	if len(columns) > 0 {
		var cols string
		for i, c := range columns {
			if i > 0 {
				cols += ", "
			}
			cols += fmt.Sprintf("AVG(%s)", c)
		}
		return newSelectStmt(conn, cols)
	}
	return newSelectStmt(conn)
}

// Max calls MAX function.
func Max(conn conn, columns ...string) iselect.Stmt {
	if len(columns) > 0 {
		var cols string
		for i, c := range columns {
			if i > 0 {
				cols += ", "
			}
			cols += fmt.Sprintf("MAX(%s)", c)
		}
		return newSelectStmt(conn, cols)
	}
	return newSelectStmt(conn)
}

// Min calls MIN function.
func Min(conn conn, columns ...string) iselect.Stmt {
	if len(columns) > 0 {
		var cols string
		for i, c := range columns {
			if i > 0 {
				cols += ", "
			}
			cols += fmt.Sprintf("MIN(%s)", c)
		}
		return newSelectStmt(conn, cols)
	}
	return newSelectStmt(conn)
}

// AlterTable calls ALTER TABLE command.
func AlterTable(conn conn, table string) ialtertable.Stmt {
	return newAlterTableStmt(conn, table)
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(conn conn, dbName string) icreatedb.Stmt {
	return newCreateDBStmt(conn, dbName)
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(conn conn, idx string) icreateindex.Stmt {
	return newCreateIndexStmt(conn, idx)
}

// CreateTable calls CREATE TABLE command.
func CreateTable(conn conn, table string) icreatetable.Stmt {
	return newCreateTableStmt(conn, table)
}

// DropDB calls DROP DATABASE command.
func DropDB(conn conn, dbName string) idropdb.Stmt {
	return newDropDBStmt(conn, dbName)
}

// DropTable calls DROP TABLE command.
func DropTable(conn conn, table string) idroptable.Stmt {
	return newDropTableStmt(conn, table)
}
