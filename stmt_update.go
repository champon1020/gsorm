package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"

	provider "github.com/champon1020/mgorm/provider/update"
)

// UpdateStmt is UPDATE statement..
type UpdateStmt struct {
	stmt
	cmd *clause.Update
}

// String returns SQL statement with string.
func (s *UpdateStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *UpdateStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Exec executes SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
func (s *UpdateStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *UpdateStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		cols := []string{}
		cols = append(cols, s.cmd.Columns...)
		if err = s.buildSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
	}

	if err = s.buildSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// buildSQLWithClauses builds SQL statement from called clauses.
func (s *UpdateStmt) buildSQLWithClauses(sql *internal.SQL) error {
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
			msg := fmt.Sprintf("%s is not supported for UPDATE statement", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *UpdateStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
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
		return errors.New("If model is not variable, model must be pointer", errors.InvalidValueError)
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
				return errors.New("Column names must be included in one of map keys", errors.InvalidSyntaxError)
			}
			vStr, err := internal.ToString(v.Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("%s = %s", c, vStr))
		}
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for (*UpdateStmt).Model", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

// Model sets model to UpdateStmt.
func (s *UpdateStmt) Model(model interface{}) provider.ModelMP {
	s.model = model
	return s
}

// Set calls SET clause.
func (s *UpdateStmt) Set(vals ...interface{}) provider.SetMP {
	if s.cmd == nil {
		s.throw(errors.New("(*UpdateStmt).cmd is nil", errors.InvalidValueError))
		return s
	}
	if len(s.cmd.Columns) != len(vals) {
		s.throw(errors.New("Number of values is not equal to that of columns", errors.InvalidValueError))
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
func (s *UpdateStmt) Where(expr string, vals ...interface{}) provider.WhereMP {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *UpdateStmt) And(expr string, vals ...interface{}) provider.AndMP {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *UpdateStmt) Or(expr string, vals ...interface{}) provider.OrMP {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}
