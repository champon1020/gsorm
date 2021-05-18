package statement

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/internal/parser"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
)

// stmt stores information about query.
type stmt struct {
	conn   domain.Conn
	called []domain.Clause
	errors []error
}

// call appends called clause.
func (s *stmt) call(e domain.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *stmt) throw(err error) {
	s.errors = append(s.errors, err)
}

// Called returns called clauses.
func (s *stmt) Called() []domain.Clause {
	return s.called
}

func (s *stmt) string(buildSQL func(*internal.SQL) error) string {
	var sql internal.SQL
	if err := buildSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

func (s *stmt) funcString(cmd domain.Clause) string {
	str := cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

func (s *stmt) compareWith(cmd domain.Clause, targetStmt domain.Stmt) error {
	expected := s.Called()
	actual := targetStmt.Called()
	if len(expected) != len(actual) {
		err := failure.New(errInvalidValue,
			failure.Context{"expected": s.funcString(cmd), "actual": targetStmt.FuncString()},
			failure.Message("statements comparison is failed"))
		return err
	}
	for i, e := range expected {
		if diff := cmp.Diff(actual[i], e); diff != "" {
			err := failure.New(errInvalidValue,
				failure.Context{"expected": s.funcString(cmd), "actual": targetStmt.FuncString()},
				failure.Message("statements comparison is failed"))
			return err
		}
	}
	return nil
}

func (s *stmt) query(buildSQL func(*internal.SQL) error, stmt domain.Stmt, model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case domain.Mock:
		returned, err := conn.CompareWith(stmt)
		if err != nil || returned == nil {
			return err
		}

		v := reflect.ValueOf(returned)
		if v.Kind() == reflect.Ptr {
			return failure.New(errInvalidValue, failure.Message("returned valud must not be pointer"))
		}
		mv := reflect.ValueOf(model)
		if mv.Kind() != reflect.Ptr {
			return failure.New(errInvalidValue, failure.Message("model must be pointer"))
		}

		mv.Elem().Set(v)
		return nil
	case domain.DB, domain.Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}

		rows, err := conn.Query(sql.String())
		if err != nil {
			return failure.Wrap(err)
		}

		defer rows.Close()
		if err := parser.MapRowsToModel(rows, model); err != nil {
			return failure.Translate(err, errFailedParse)
		}
		return nil
	}

	return failure.New(errInvalidValue,
		failure.Context{"conn": reflect.TypeOf(s.conn).String()},
		failure.Message("conn can be *DB, *Tx, *MockDB or *MockTx"))
}

func (s *stmt) exec(buildSQL func(*internal.SQL) error, stmt domain.Stmt) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case domain.Mock:
		_, err := conn.CompareWith(stmt)
		if err != nil {
			return err
		}
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
		failure.Message("type of conn can be *DB, *Tx, *MockDB or *MockTx"))
}
