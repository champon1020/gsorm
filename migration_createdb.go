package mgorm

import (
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
)

// CreateDBStmt is CREATE DATABASE statement.
type CreateDBStmt struct {
	migStmt
	cmd *mig.CreateDB
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
