package mgorm

import (
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
)

// DropTableStmt is DROP TABLE statement.
type DropTableStmt struct {
	migStmt
	cmd *mig.DropTable
}

func (s *DropTableStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *DropTableStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *DropTableStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}
