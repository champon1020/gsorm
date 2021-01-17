package mgorm

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/champon1020/mgorm/internal"
)

// Op values for error handling.
const (
	OpSQLDoQuery internal.Op = "mgorm.SQL.doQuery"
	OpSQLDoExec  internal.Op = "mgorm.SQL.doExec"
	OpSetField   internal.Op = "mgorm.setField"
)

// SQL string.
type SQL string

func (s *SQL) string() string {
	return string(*s)
}

func (s *SQL) write(str string) {
	if len(*s) != 0 && str != ")" {
		*s += " "
	}
	*s += SQL(str)
}

func (s *SQL) doQuery(db internal.DB, model interface{}) error {
	rows, err := db.Query(s.string())
	if err != nil {
		return internal.NewError(OpSQLDoQuery, internal.KindDatabase, err)
	}
	if rows == nil {
		return internal.NewError(OpSQLDoQuery, internal.KindDatabase, errors.New("rows is nil"))
	}

	cols, err := rows.Columns()
	if err != nil {
		return internal.NewError(OpSQLDoQuery, internal.KindDatabase, err)
	}

	rowVal := make([]interface{}, len(cols))
	rowValPtr := []interface{}{}
	for i := 0; i < len(rowVal); i++ {
		rowValPtr = append(rowValPtr, &rowVal[i])
	}

	mt := reflect.TypeOf(model).Elem()
	mv := reflect.New(mt).Elem()

	// Model type must be slice or array.
	if mt == nil || (mt.Kind() != reflect.Slice && mt.Kind() != reflect.Array) {
		err := errors.New("model type must be slice or array")
		return internal.NewError(OpSQLDoQuery, internal.KindType, err)
	}

	for rows.Next() {
		if err := rows.Scan(rowValPtr...); err != nil {
			return internal.NewError(OpSQLDoQuery, internal.KindDatabase, err)
		}

		if err := setToModel(&mv, mt, cols, rowVal); err != nil {
			return err
		}
	}

	rows.Close()

	modelRef := reflect.ValueOf(model).Elem()
	modelRef.Set(mv)

	return nil
}

func (s *SQL) doExec(db internal.DB) error {
	_, err := db.Exec(s.string())
	if err != nil {
		return internal.NewError(OpSQLDoExec, internal.KindDatabase, err)
	}
	return nil
}

func setToModel(mv *reflect.Value, mt reflect.Type, cols []string, rowVal []interface{}) error {
	// Generate reflect type and value for model.
	t := mt.Elem()
	v := reflect.New(t)

	// Loop with columns of rows.
	for i, c := range cols {
		for j := 0; j < t.NumField(); j++ {
			// Check column name.
			if c != columnName(t.Field(j)) {
				continue
			}

			// Set values to struct fields.
			if err := setField(reflect.Indirect(v).Field(j), rowVal[i]); err != nil {
				return err
			}
			break
		}
	}

	// Append struct to slice (or array).
	*mv = reflect.Append(*mv, v.Elem())
	return nil
}

func columnName(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return internal.ConvertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}

func setField(f reflect.Value, v interface{}) error {
	if !f.CanSet() {
		err := errors.New("field cannot be changes")
		return internal.NewError(OpSetField, internal.KindBasic, err)
	}

	switch f.Kind() {
	case reflect.String:
		sv, ok := v.(string)
		if !ok {
			err := errors.New("field type is invalid")
			return internal.NewError(OpSetField, internal.KindType, err)
		}
		f.SetString(sv)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src := fmt.Sprintf("%v", v)
		i64, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			err := errors.New("field type is invalid")
			return internal.NewError(OpSetField, internal.KindType, err)
		}
		f.SetInt(i64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src := fmt.Sprintf("%v", v)
		u64, err := strconv.ParseUint(src, 10, 64)
		if err != nil {
			err := errors.New("field type is invalid")
			return internal.NewError(OpSetField, internal.KindType, err)

		}
		f.SetUint(u64)
	case reflect.Float32, reflect.Float64:
		src := fmt.Sprintf("%v", v)
		f64, err := strconv.ParseFloat(src, 64)
		if err != nil {
			err := errors.New("field type is invalid")
			return internal.NewError(OpSetField, internal.KindType, err)

		}
		f.SetFloat(f64)
	case reflect.Bool:
		b, ok := v.(bool)
		if !ok {
			err := errors.New("field type is invalid")
			return internal.NewError(OpSetField, internal.KindType, err)

		}
		f.SetBool(b)
	case reflect.Struct:
		if f.Type() == reflect.TypeOf(time.Time{}) {
			t, ok := v.(time.Time)
			if !ok {
				err := errors.New("field type is invalid")
				return internal.NewError(OpSetField, internal.KindType, err)

			}
			f.Set(reflect.ValueOf(t))
		}
	}

	return nil
}
