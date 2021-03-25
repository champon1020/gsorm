package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"

	ifc "github.com/champon1020/mgorm/interfaces/delete"
)

// DeleteStmt is DELETE statement.
type DeleteStmt struct {
	stmt
	cmd *clause.Delete
}

// String returns SQL statement with string.
func (s *DeleteStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *DeleteStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Cmd returns cmd clause.
func (s *DeleteStmt) Cmd() syntax.Clause {
	return s.cmd
}

// Exec executed SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
func (s *DeleteStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *DeleteStmt) buildSQL(sql *internal.SQL) error {
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
func (s *DeleteStmt) From(tables ...string) ifc.From {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Where calls WHERE clause.
func (s *DeleteStmt) Where(expr string, vals ...interface{}) ifc.Where {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *DeleteStmt) And(expr string, vals ...interface{}) ifc.And {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *DeleteStmt) Or(expr string, vals ...interface{}) ifc.Or {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}