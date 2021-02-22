package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
)

type MgormDelete interface {
	From(...string) DeleteFrom
}

type DeleteFrom interface {
	Where(string, ...interface{}) DeleteWhere
	ExpectExec() *DeleteStmt
	ExecCallable
}

type DeleteWhere interface {
	And(string, ...interface{}) DeleteAnd
	Or(string, ...interface{}) DeleteOr
	ExpectExec() *DeleteStmt
	ExecCallable
}

type DeleteAnd interface {
	ExpectExec() *DeleteStmt
	ExecCallable
}

type DeleteOr interface {
	ExpectExec() *DeleteStmt
	ExecCallable
}

// DeleteStmt is DELETE statement.
type DeleteStmt struct {
	stmt
	cmd *clause.Delete
}

func (s *DeleteStmt) String() string {
	var sql internal.SQL
	if err := s.processSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

func (s *DeleteStmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

func (s *DeleteStmt) ExpectExec() *DeleteStmt {
	return s
}

func (s *DeleteStmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB, *Tx:
		var sql internal.SQL
		ss, err := s.cmd.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())

		if err := s.processSQL(&sql); err != nil {
			return err
		}
		if _, err := pool.Exec(sql.String()); err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}
		return nil
	case Mock:
		_, err := pool.CompareWith(s)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("DB type must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
}

func (s *DeleteStmt) processSQL(sql *internal.SQL) error {
	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.From:
			s, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("Type %s is not supported for DELETE", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}
	return nil
}

// From calls FROM clause.
func (s *DeleteStmt) From(tables ...string) DeleteFrom {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Where calls WHERE clause.
func (s *DeleteStmt) Where(expr string, vals ...interface{}) DeleteWhere {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *DeleteStmt) And(expr string, vals ...interface{}) DeleteAnd {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *DeleteStmt) Or(expr string, vals ...interface{}) DeleteOr {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}
