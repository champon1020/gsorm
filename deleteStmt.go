package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
)

// MgormDelete is interface for returned value of mgorm.Delete.
type MgormDelete interface {
	From(...string) DeleteFrom
}

// DeleteFrom is interface for returned value of (*DeleteStmt).From.
type DeleteFrom interface {
	Where(string, ...interface{}) DeleteWhere
	ExpectExec() *DeleteStmt
	ExecCallable
}

// DeleteWhere is interface for returned value of (*DeleteStmt).Where.
type DeleteWhere interface {
	And(string, ...interface{}) DeleteAnd
	Or(string, ...interface{}) DeleteOr
	ExpectExec() *DeleteStmt
	ExecCallable
}

// DeleteAnd is interface for returned value of (*DeleteStmt).And.
type DeleteAnd interface {
	ExpectExec() *DeleteStmt
	ExecCallable
}

// DeleteOr is interface for returned value of (*DeleteOr).Or.
type DeleteOr interface {
	ExpectExec() *DeleteStmt
	ExecCallable
}

// DeleteStmt is DELETE statement.
type DeleteStmt struct {
	stmt
	cmd *clause.Delete
}

// String returns SQL statement with string.
func (s *DeleteStmt) String() string {
	var sql internal.SQL
	if err := s.processSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

// funcString returns function call as string.
func (s *DeleteStmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

// ExpectExec returns *DeleteStmt. This function is used for mock test.
func (s *DeleteStmt) ExpectExec() *DeleteStmt {
	return s
}

// Exec executed SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
func (s *DeleteStmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case *DB, *Tx:
		var sql internal.SQL
		if err := s.processSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}
		return nil
	case Mock:
		_, err := conn.CompareWith(s)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Type of conn must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
}

// processSQL builds SQL statement.
func (s *DeleteStmt) processSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.From,
			*clause.Where,
			*clause.And,
			*clause.Or:
			s, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("%s is not supported for DELETE statement", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidSyntaxError)
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
