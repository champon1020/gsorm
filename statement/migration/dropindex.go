package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"

	ifc "github.com/champon1020/mgorm/interfaces/dropindex"
)

// DropIndexStmt is DROP INDEX statement.
type DropIndexStmt struct {
	migStmt
	cmd *mig.DropIndex
}

func NewDropIndexStmt(conn domain.Conn, idx string) *DropIndexStmt {
	stmt := &DropIndexStmt{cmd: &mig.DropIndex{IdxName: idx}}
	stmt.conn = conn
	return stmt
}

func (s *DropIndexStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *DropIndexStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *DropIndexStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			break
		}

		switch e := e.(type) {
		case *mig.On:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).String()},
				failure.Message("invalid clause for DROP INDEX"))
		}
	}

	return nil
}

// On calls ON clause.
func (s *DropIndexStmt) On(table string) ifc.On {
	s.call(&mig.On{Table: table})
	return s
}
