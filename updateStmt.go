package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
)

type MgormUpdate interface {
	Model(interface{}) UpdateModel
	Set(...interface{}) UpdateSet
}

type UpdateModel interface {
	Where(string, ...interface{}) UpdateWhere
	ExpectExec() *UpdateStmt
	ExecCallable
}

type UpdateSet interface {
	Where(string, ...interface{}) UpdateWhere
	ExpectExec() *UpdateStmt
	ExecCallable
}

type UpdateWhere interface {
	And(string, ...interface{}) UpdateAnd
	Or(string, ...interface{}) UpdateOr
	ExpectExec() *UpdateStmt
	ExecCallable
}

type UpdateAnd interface {
	ExpectExec() *UpdateStmt
	ExecCallable
}

type UpdateOr interface {
	ExpectExec() *UpdateStmt
	ExecCallable
}

// UpdateStmt is UPDATE STATEMENT.
type UpdateStmt struct {
	stmt
	cmd *clause.Update
}

func (s *UpdateStmt) String() string {
	var sql internal.SQL
	if err := s.processSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

func (s *UpdateStmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

func (s *UpdateStmt) ExpectExec() *UpdateStmt {
	return s
}

func (s *UpdateStmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB, *Tx:
		var sql internal.SQL
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

func (s *UpdateStmt) processSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		cols := []string{}
		for _, c := range s.cmd.Columns {
			cols = append(cols, c)
		}
		if err := s.processSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
	}

	if err = s.processSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

func (s *UpdateStmt) processSQLWithClauses(sql *internal.SQL) error {
	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.Set,
			*clause.Where,
			*clause.And,
			*clause.Or:
			s, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("Type %s is not supported for UPDATE", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}
	return nil
}

func (s *UpdateStmt) processSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	ref := reflect.ValueOf(model)
	switch ref.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool,
		reflect.String:
		if len(cols) != 1 {
			msg := fmt.Sprintf("If you set variable to Model, number of columns must be 1, not %d", len(cols))
			return errors.New(msg, errors.InvalidSyntaxError)
		}
		vStr, err := internal.ToString(ref.Interface(), true)
		if err != nil {
			return err
		}
		sql.Write(fmt.Sprintf("SET %s = %s", cols[0], vStr))
		return nil
	}

	if ref.Kind() != reflect.Ptr {
		return errors.New("Model must be pointer", errors.InvalidValueError)
	}
	ref = ref.Elem()

	sql.Write("SET")
	switch ref.Kind() {
	case reflect.Struct:
		idxC2F := internal.MapOfColumnsToFields(cols, reflect.TypeOf(ref.Interface()))
		for i, c := range cols {
			if i > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Field(idxC2F[i]).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("%s = %s", c, vStr))
		}
		return nil
	case reflect.Map:
		for i, c := range cols {
			if i > 0 {
				sql.Write(",")
			}
			v := ref.MapIndex(reflect.ValueOf(c))
			if !v.IsValid() {
				return errors.New("Column names must be included in some of map keys", errors.InvalidSyntaxError)
			}
			vStr, err := internal.ToString(v.Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("%s = %s", c, vStr))
		}
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for Model with UPDATE", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

// Model sets model to Stmt.
func (s *UpdateStmt) Model(model interface{}) UpdateModel {
	s.model = model
	return s
}

// Set calls SET clause.
func (s *UpdateStmt) Set(vals ...interface{}) UpdateSet {
	if s.cmd == nil {
		s.throw(errors.New("Command is nil", errors.InvalidValueError))
		return s
	}
	if len(s.cmd.Columns) != len(vals) {
		s.throw(errors.New("Length is different between columns and values", errors.InvalidValueError))
		return s
	}
	set := new(clause.Set)
	for i, c := range s.cmd.Columns {
		set.AddEq(c, vals[i])
	}
	s.call(set)
	return s
}

// Where calls WHERE clause.
func (s *UpdateStmt) Where(expr string, vals ...interface{}) UpdateWhere {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *UpdateStmt) And(expr string, vals ...interface{}) UpdateAnd {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *UpdateStmt) Or(expr string, vals ...interface{}) UpdateOr {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}
