package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"

	ifc "github.com/champon1020/mgorm/interfaces/createindex"
)

// CreateIndexStmt is CREATE INDEX statement.
type CreateIndexStmt struct {
	migStmt
	cmd *mig.CreateIndex
}

func (s *CreateIndexStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *CreateIndexStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *CreateIndexStmt) buildSQL(sql *internal.SQL) error {
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
			msg := fmt.Sprintf("%v is not supported for CREATE INDEX statement", reflect.TypeOf(e).String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}

	return nil
}

// On calls ON clause.
func (s *CreateIndexStmt) On(table string, cols ...string) ifc.On {
	s.call(&mig.On{Table: table, Columns: cols})
	return s
}
