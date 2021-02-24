package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
)

// MgormInsert is interface for returned value of mgorm.Insert.
type MgormInsert interface {
	Model(interface{}) InsertModel
	Values(...interface{}) InsertValues
}

// InsertModel is interface for returned value of (*InsertStmt).Model.
type InsertModel interface {
	ExpectExec() *InsertStmt
	ExecCallable
}

// InsertValues is interface for returned value of (*InsertStmt).Values.
type InsertValues interface {
	Values(...interface{}) InsertValues
	ExpectExec() *InsertStmt
	ExecCallable
}

// InsertStmt is INSERT statement.
type InsertStmt struct {
	stmt
	cmd *clause.Insert
}

// String returns SQL statement with string.
func (s *InsertStmt) String() string {
	var sql internal.SQL
	if err := s.processSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

// funcString returns function call as string.
func (s *InsertStmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

// ExpectExec returns *InsertStmt. This function is used for mock test.
func (s *InsertStmt) ExpectExec() *InsertStmt {
	return s
}

// Exec executed SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
func (s *InsertStmt) Exec() error {
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

// processSQL builds SQL statement.
func (s *InsertStmt) processSQL(sql *internal.SQL) error {
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
		if err := s.processSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
		return nil
	}

	if err := s.processSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// processSQLWithClauses builds SQL statement from called clauses.
func (s *InsertStmt) processSQLWithClauses(sql *internal.SQL) error {
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
			msg := fmt.Sprintf("Type %s is not supported for INSERT", reflect.TypeOf(e).Elem().String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}
	return nil
}

// processSQLWithModel builds SQL statement from model.
func (s *InsertStmt) processSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
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
				return errors.New("Column names must be included in some of map keys", errors.InvalidSyntaxError)
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

	msg := fmt.Sprintf("Type %s is not supported for Model with INSERT", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

// Model sets model to InsertStmt.
func (s *InsertStmt) Model(model interface{}) InsertModel {
	s.model = model
	return s
}

// Values calls VALUES clause.
func (s *InsertStmt) Values(vals ...interface{}) InsertValues {
	v := new(clause.Values)
	for _, val := range vals {
		v.AddValue(val)
	}
	s.call(v)
	return s
}
