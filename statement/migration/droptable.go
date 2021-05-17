package migration

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
)

// DropTableStmt is DROP TABLE statement.
type DropTableStmt struct {
	migStmt
	cmd *mig.DropTable
}

// NewDropTableStmt creates DropTableStmt instance.
func NewDropTableStmt(conn domain.Conn, table string) *DropTableStmt {
	stmt := &DropTableStmt{cmd: &mig.DropTable{Table: table}}
	stmt.conn = conn
	return stmt
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
