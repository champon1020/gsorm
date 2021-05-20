package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/interfaces/idroptable"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"
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

	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			break
		}

		switch e := e.(type) {
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).String()},
				failure.Message("invalid clause for CREATE DATABASE"))
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *DropTableStmt) RawClause(rs string, v ...interface{}) idroptable.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}
