package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
)

// migStmt stores information about database migration query.
type migStmt struct {
	conn   Conn
	called []syntax.MigClause
	errors []error
}

// call appends called clause.
func (s *migStmt) call(e syntax.MigClause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *migStmt) throw(e error) {
	s.errors = append(s.errors, e)
}

// headClause returns first element of called.
func (s *migStmt) headClause() syntax.MigClause {
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
	case *DB, *Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}
	case *MockDB, *MockTx:
		return nil
	}

	return errors.New("Type of conn must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
}

func (s *migStmt) buildColumnOptSQL(sql *internal.SQL) error {
	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			return nil
		}

		switch e := e.(type) {
		case *mig.NotNull,
			*mig.Default:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		case *mig.AutoInc:
			if s.conn.getDriver() == internal.PSQL {
				return errors.New("AUTO_INCREMENT clause is not allowed in PostgreSQL", errors.InvalidSyntaxError)
			}
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
		msg := "Called claues have already been processed but SQL statement is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
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

	msg := fmt.Sprintf("%v is not supported for CONSTRAINT statement", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (s *migStmt) buildRefSQL(sql *internal.SQL) error {
	e := s.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL statement is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
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

	msg := fmt.Sprintf("%v is not supported for CONSTRAINT FOREIGN KEY statement", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}
