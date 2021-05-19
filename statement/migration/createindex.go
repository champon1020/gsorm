package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/interfaces/icreateindex"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"
)

// CreateIndexStmt is CREATE INDEX statement.
type CreateIndexStmt struct {
	migStmt
	cmd *mig.CreateIndex
}

// NewCreateIndexStmt creates CreateIndexStmt instance.
func NewCreateIndexStmt(conn domain.Conn, idx string) *CreateIndexStmt {
	stmt := &CreateIndexStmt{cmd: &mig.CreateIndex{IdxName: idx}}
	stmt.conn = conn
	return stmt
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
		case *syntax.RawClause,
			*mig.On:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).String()},
				failure.Message("invalid clause for CREATE INDEX"))
		}
	}

	return nil
}

// On calls ON clause.
func (s *CreateIndexStmt) On(table string, cols ...string) icreateindex.On {
	s.call(&mig.On{Table: table, Columns: cols})
	return s
}

// RawClause calls the raw string clause.
func (s *CreateIndexStmt) RawClause(rs string, v ...interface{}) icreateindex.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}
