package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"

	provider "github.com/champon1020/mgorm/provider/drop"
)

// DropDBStmt is DROP DATABASE statement.
type DropDBStmt struct {
	migStmt
	cmd *mig.DropDB
}

func (s *DropDBStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *DropDBStmt) Migration() error {
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

// DropIndexStmt is DROP INDEX statement.
type DropIndexStmt struct {
	migStmt
	cmd *mig.DropIndex
}

func (s *DropIndexStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *DropIndexStmt) Migration() error {
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
			if s.conn.getDriver() != internal.MySQL {
				msg := "DROP INDEX command with ON clause is not allowed in PostgreSQL"
				return errors.New(msg, errors.InvalidSyntaxError)
			}
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			msg := fmt.Sprintf("%v is not supported for DROP INDEX statement", reflect.TypeOf(e).String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}

	return nil
}

// On calls ON clause.
func (s *DropIndexStmt) On(table string) provider.OnMP {
	s.call(&mig.On{Table: table})
	return s
}

// DropTableStmt is DROP TABLE statement.
type DropTableStmt struct {
	migStmt
	cmd *mig.DropTable
}

func (s *DropTableStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *DropTableStmt) Migration() error {
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
