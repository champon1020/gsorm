package mgorm

import (
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
)

// DropDBStmt is DROP DATABASE statement.
type DropDBStmt struct {
	migStmt
	cmd *mig.DropDB
}

func (s *DropDBStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *DropDBStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *DropDBStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}
