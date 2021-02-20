package mgorm

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
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
func Select(db Pool, cols ...string) SelectStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = clause.NewSelect(cols)
	return stmt
}

// Insert calls INSERT command.
func Insert(db Pool, table string, cols ...string) InsertStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = clause.NewInsert(table, cols)
	return stmt
}

// Update calls UPDATE command.
func Update(db Pool, table string, cols ...string) UpdateStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = clause.NewUpdate(table, cols)
	return stmt
}

// Delete calls DELETE command.
func Delete(db Pool) DeleteStmt {
	stmt := &Stmt{db: db}
	stmt.cmd = clause.NewDelete()
	return stmt
}

// Count calls COUNT function.
func Count(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = clause.NewSelect([]string{s})
	return stmt
}

// Avg calls AVG function.
func Avg(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = clause.NewSelect([]string{s})
	return stmt
}

// Sum calls SUM function.
func Sum(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = clause.NewSelect([]string{s})
	return stmt
}

// Min calls MIN function.
func Min(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = clause.NewSelect([]string{s})
	return stmt
}

// Max calls MAX function.
func Max(db Pool, col string, alias ...string) SelectStmt {
	stmt := &Stmt{db: db}
	s := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		s = fmt.Sprintf("%s AS %s", s, alias[0])
	}
	stmt.cmd = clause.NewSelect([]string{s})
	return stmt
}

// CreateDB calls CREATE DATABASE command.
func CreateDB(pool Pool, dbName string) CreateDBMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.CreateDB{DBName: dbName},
	}
}

// DropDB calls DROP DATABASE command.
func DropDB(pool Pool, dbName string) DropDBMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.DropDB{DBName: dbName},
	}
}

// CreateTable calls CREATE TABLE command.
func CreateTable(pool Pool, table string) CreateTableMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.CreateTable{Table: table},
	}
}

// DropTable calls DROP TABLE command.
func DropTable(pool Pool, table string) DropTableMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.DropTable{Table: table},
	}
}

// AlterTable calls ALTER TABLE command.
func AlterTable(pool Pool, table string) AlterTableMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.AlterTable{Table: table},
	}
}

// CreateIndex calls CREATE INDEX command.
func CreateIndex(pool Pool, idx string) CreateIndexMig {
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.CreateIndex{IdxName: idx},
	}
}

// DropIndex calls DROP INDEX command.
func DropIndex(pool Pool, table string, idx string) DropIndexMig {
	if pool.getDriver() == internal.MySQL {
		return &MigStmt{
			pool:   pool,
			driver: pool.getDriver(),
			cmd:    &mig.AlterTable{Table: table},
			called: []syntax.MigClause{&mig.DropIndex{IdxName: idx}},
		}
	}
	return &MigStmt{
		pool:   pool,
		driver: pool.getDriver(),
		cmd:    &mig.DropIndex{IdxName: idx},
	}
}
