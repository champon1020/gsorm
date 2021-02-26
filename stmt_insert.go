package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/provider"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"

	prInsert "github.com/champon1020/mgorm/provider/insert"
)

// InsertStmt is INSERT statement.
type InsertStmt struct {
	stmt
	model interface{}
	cmd   *clause.Insert
	sel   Stmt
}

// String returns SQL statement with string.
func (s *InsertStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *InsertStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Cmd returns cmd clause.
func (s *InsertStmt) Cmd() syntax.Clause {
	return s.cmd
}

// Exec executed SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
func (s *InsertStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *InsertStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		cols := []string{}
		for _, c := range s.cmd.Columns {
			if c.Alias != "" {
				cols = append(cols, c.Alias)
				continue
			}
			cols = append(cols, c.Name)
		}
		if err := s.buildSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
		return nil
	}

	if s.sel != nil {
		sql.Write(s.sel.String())
		return nil
	}

	if err := s.buildSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// buildSQLWithClauses builds SQL statement from called clauses.
func (s *InsertStmt) buildSQLWithClauses(sql *internal.SQL) error {
	for i, e := range s.called {
		switch e := e.(type) {
		case *clause.Values:
			s, err := e.Build()
			if err != nil {
				return err
			}
			if i > 0 {
				sql.Write(",")
				sql.Write(s.BuildValue())
				continue
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("%s is not supported for INSERT statement", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidSyntaxError)
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *InsertStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	ref := reflect.ValueOf(model)
	if ref.Kind() != reflect.Ptr {
		return errors.New("Model must be pointer", errors.InvalidValueError)
	}
	ref = ref.Elem()

	sql.Write("VALUES")
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		// Type of slice element.
		typ := reflect.TypeOf(ref.Interface()).Elem()

		// If undelying type of slice element is struct.
		if typ.Kind() == reflect.Struct {
			idxC2F := internal.MapOfColumnsToFields(cols, typ)
			for i := 0; i < ref.Len(); i++ {
				if i > 0 {
					sql.Write(",")
				}
				sql.Write("(")
				for j := 0; j < len(cols); j++ {
					if j > 0 {
						sql.Write(",")
					}
					vStr, err := internal.ToString(ref.Index(i).Field(idxC2F[j]).Interface(), true)
					if err != nil {
						return err
					}
					sql.Write(vStr)
				}
				sql.Write(")")
			}
			return nil
		}

		for i := 0; i < ref.Len(); i++ {
			if i > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Index(i).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("(%s)", vStr))
		}
		return nil
	case reflect.Struct:
		idxC2F := internal.MapOfColumnsToFields(cols, reflect.TypeOf(ref.Interface()))
		sql.Write("(")
		for j := 0; j < len(cols); j++ {
			if j > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Field(idxC2F[j]).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(vStr)
		}
		sql.Write(")")
		return nil
	case reflect.Map:
		sql.Write("(")
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
			sql.Write(vStr)
		}
		sql.Write(")")
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for (*InsertStmt).Model", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

// Model sets model to InsertStmt.
func (s *InsertStmt) Model(model interface{}) prInsert.ModelMP {
	s.model = model
	return s
}

// Select calls SELECT statement.
func (s *InsertStmt) Select(sel provider.QueryCallable) prInsert.SelectMP {
	s.sel = sel
	return s
}

// Values calls VALUES clause.
func (s *InsertStmt) Values(vals ...interface{}) prInsert.ValuesMP {
	v := new(clause.Values)
	for _, val := range vals {
		v.AddValue(val)
	}
	s.call(v)
	return s
}
