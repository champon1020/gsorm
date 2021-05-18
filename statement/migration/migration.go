package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"
)

// migStmt stores information about database migration query.
type migStmt struct {
	conn   domain.Conn
	called []domain.Clause
	errors []error
}

// call appends called clause.
func (s *migStmt) call(e domain.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *migStmt) throw(e error) {
	s.errors = append(s.errors, e)
}

// headClause returns first element of called.
func (s *migStmt) headClause() domain.Clause {
	if len(s.called) == 0 {
		return nil
	}
	return s.called[0]
}

// advanceClause slides slice of called.
func (s *migStmt) advanceClause() {
	s.called = s.called[1:]
}

func (s *migStmt) string(buildSQL func(*internal.SQL) error) string {
	var sql internal.SQL
	if err := buildSQL(&sql); err != nil {
		s.throw(err)
		return ""
	}
	return sql.String()
}

func (s *migStmt) migration(buildSQL func(*internal.SQL) error) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case domain.Mock:
		return nil
	case domain.DB, domain.Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return failure.Wrap(err)
		}
		return nil
	}

	return failure.New(errInvalidValue,
		failure.Context{"conn": reflect.TypeOf(s.conn).String()},
		failure.Message("conn can be *DB, *Tx, *MockDB or *MockTx"))
}

func (s *migStmt) buildColumnOptSQL(sql *internal.SQL) error {
	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			return nil
		}

		switch e := e.(type) {
		case *syntax.RawClause,
			*mig.NotNull,
			*mig.Default:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		default:
			return nil
		}

		s.advanceClause()
	}

	return nil
}

func (s *migStmt) buildConstraintSQL(sql *internal.SQL) error {
	e := s.headClause()
	if e == nil {
		return failure.New(errInvalidSyntax,
			failure.Message("the SQL statement is not completed or the syntax is not supported"))
	}

	if rc, ok := e.(*syntax.RawClause); ok {
		ss, err := rc.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		e = s.headClause()
	}

	switch e := e.(type) {
	case *mig.Primary, *mig.Unique:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return nil
	case *mig.Foreign:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return s.buildRefSQL(sql)
	}

	return failure.New(errInvalidClause,
		failure.Context{"clause": reflect.TypeOf(e).String()},
		failure.Message("invalid clause for CONSTRAINT"))
}

func (s *migStmt) buildRefSQL(sql *internal.SQL) error {
	e := s.headClause()
	if e == nil {
		return failure.New(errInvalidSyntax,
			failure.Message("the SQL statement is not completed or the syntax is not supported"))
	}

	if rc, ok := e.(*syntax.RawClause); ok {
		ss, err := rc.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		e = s.headClause()
	}

	switch e := e.(type) {
	case *mig.Ref:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return nil
	}

	return failure.New(errInvalidClause,
		failure.Context{"clause": reflect.TypeOf(e).String()},
		failure.Message("invalid clause for FOREIGN KEY"))
}
