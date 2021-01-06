package syntax

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time"
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

func (s *SQL) doQuery(db *sql.DB, model interface{}) error {
	rows, err := db.Query(s.string())
	if err != nil || rows == nil {
		return newError(ErrQueryFailed, "database query failed")
	}

	cols, err := rows.Columns()
	if err != nil {
		return newError(ErrUnknown, "it cannot get columns from rows")
	}

	rowVal := make([]interface{}, len(cols))
	rowValPtr := []interface{}{}
	for i := 0; i < len(rowVal); i++ {
		rowValPtr = append(rowValPtr, &rowVal[i])
	}

	mt := reflect.TypeOf(model).Elem()
	mv := reflect.New(mt)

	// Model type must be slice or array.
	if mt == nil || (mt.Kind() != reflect.Slice && mt.Kind() != reflect.Array) {
		return newError(ErrInvalidType, "model type must be slice or array")
	}

	for rows.Next() {
		if err := rows.Scan(rowValPtr...); err != nil {
			return newError(ErrScanFailed, "database scan failed")
		}

		if err := setToModel(mv, mt, cols, rowVal); err != nil {
			return err
		}
	}

	modelRef := reflect.ValueOf(model).Elem()
	modelRef.Set(mv)

	return nil
}

func (s *SQL) doExec(db *sql.DB) error {
	_, err := db.Exec(s.string())
	if err != nil {
		return newError(ErrExecFailed, "database exec failed")
	}
	return nil
}

func setToModel(mv reflect.Value, mt reflect.Type, cols []string, rowVal []interface{}) error {
	// Generate reflect type and value for model struct.
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
	mv = reflect.Append(mv, v.Elem())
	return nil
}

func columnName(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return convertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}

func setField(f reflect.Value, v interface{}) error {
	if !f.CanSet() {
		return newError(ErrInvalid, "field cannot be changed")
	}

	switch f.Kind() {
	case reflect.String:
		sv, ok := v.(string)
		if !ok {
			return newError(ErrInvalidType, "field type is invalid")
		}
		f.SetString(sv)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src := fmt.Sprintf("%v", v)
		i64, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return newError(ErrInvalidType, "field type is invalid")
		}
		f.SetInt(i64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src := fmt.Sprintf("%v", v)
		u64, err := strconv.ParseUint(src, 10, 64)
		if err != nil {
			return newError(ErrInvalidType, "field type is invalid")
		}
		f.SetUint(u64)
	case reflect.Float32, reflect.Float64:
		src := fmt.Sprintf("%v", v)
		f64, err := strconv.ParseFloat(src, 64)
		if err != nil {
			return newError(ErrInvalidType, "field type is invalid")
		}
		f.SetFloat(f64)
	case reflect.Bool:
		b, ok := v.(bool)
		if !ok {
			return newError(ErrInvalidType, "field type is invalid")
		}
		f.SetBool(b)
	case reflect.Struct:
		if f.Type() == reflect.TypeOf(time.Time{}) {
			t, ok := v.(time.Time)
			if !ok {
				return newError(ErrInvalidType, "field type is invalid")
			}
			f.Set(reflect.ValueOf(t))
		}
	}

	return nil
}
