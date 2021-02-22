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
func Select(p Pool, cols ...string) MgormSelect {
	s := new(SelectStmt)
	s.db = p
	s.cmd = clause.NewSelect(cols)
	return s
}

// Insert calls INSERT command.
func Insert(p Pool, table string, cols ...string) MgormInsert {
	s := new(InsertStmt)
	s.db = p
	s.cmd = clause.NewInsert(table, cols)
	return s
}

// Update calls UPDATE command.
func Update(p Pool, table string, cols ...string) MgormUpdate {
	s := new(UpdateStmt)
	s.db = p
	s.cmd = clause.NewUpdate(table, cols)
	return s
}

// Delete calls DELETE command.
func Delete(p Pool) MgormDelete {
	s := new(DeleteStmt)
	s.db = p
	s.cmd = clause.NewDelete()
	return s
}

// Count calls COUNT function.
func Count(p Pool, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("COUNT(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.db = p
	return s
}

// Avg calls AVG function.
func Avg(p Pool, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("AVG(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.db = p
	return s
}

// Sum calls SUM function.
func Sum(p Pool, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("SUM(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.db = p
	return s
}

// Min calls MIN function.
func Min(p Pool, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("MIN(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.db = p
	return s
}

// Max calls MAX function.
func Max(p Pool, col string, alias ...string) MgormSelect {
	c := fmt.Sprintf("MAX(%s)", col)
	if len(alias) > 0 {
		c = fmt.Sprintf("%s AS %s", c, alias[0])
	}
	s := &SelectStmt{cmd: clause.NewSelect([]string{c})}
	s.db = p
	return s
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
