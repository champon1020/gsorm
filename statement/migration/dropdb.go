package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/interfaces/idropdb"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"
)

// DropDBStmt is DROP DATABASE statement.
type DropDBStmt struct {
	migStmt
	cmd *mig.DropDB
}

// NewDropDBStmt creates DropDBStmt instance.
func NewDropDBStmt(conn domain.Conn, dbName string) *DropDBStmt {
	stmt := &DropDBStmt{cmd: &mig.DropDB{DBName: dbName}}
	stmt.conn = conn
	return stmt
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
func (s *DropDBStmt) RawClause(rs string, v ...interface{}) idropdb.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}
