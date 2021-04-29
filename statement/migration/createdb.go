package migration

import (
	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
)

// CreateDBStmt is CREATE DATABASE statement.
type CreateDBStmt struct {
	migStmt
	cmd *mig.CreateDB
}

func NewCreateDBStmt(conn database.Conn, dbName string) *CreateDBStmt {
	stmt := &CreateDBStmt{cmd: &mig.CreateDB{DBName: dbName}}
	stmt.conn = conn
	return stmt
}

func (s *CreateDBStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *CreateDBStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *CreateDBStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}
