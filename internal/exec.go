package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Query executes query and sets rows to model structure.
func Query(db *sql.DB, s *SQL, model interface{}) error {
	// Execute query.
	rows, err := db.Query(s.String())
	if err != nil {
		return NewError(opSQLDoQuery, KindDatabase, err)
	}
	if rows == nil {
		return NewError(opSQLDoQuery, KindDatabase, errors.New("rows is nil"))
	}

	// Model reflection.
	mt := reflect.TypeOf(model).Elem()
	mv := reflect.New(mt).Elem()

	// Model type must be slice or array.
	if mt == nil || (mt.Kind() != reflect.Slice && mt.Kind() != reflect.Array) {
		err := errors.New("model type must be slice or array")
		return NewError(opSQLDoQuery, KindType, err)
	}

	// Generate map to localize field index between row and model.
	rCols, err := rows.Columns()
	if err != nil {
		return NewError(opSQLDoQuery, KindDatabase, err)
	}
	indR2M := make(map[int]int)
	for i, c := range rCols {
		for j := 0; j < mt.Elem().NumField(); j++ {
			if c != columnName(mt.Elem().Field(j)) {
				continue
			}
			indR2M[i] = j
		}
	}

	// Prepare pointers which is used to rows.Scan().
	rVal := make([][]byte, len(rCols))
	rValPtr := []interface{}{}
	for i := 0; i < len(rVal); i++ {
		rValPtr = append(rValPtr, &rVal[i])
	}

	for rows.Next() {
		if err := rows.Scan(rValPtr...); err != nil {
			return NewError(opSQLDoQuery, KindDatabase, err)
		}
		if err := setToModel(&mv, mt, &indR2M, rVal); err != nil {
			return err
		}
	}
	rows.Close()

	modelRef := reflect.ValueOf(model).Elem()
	modelRef.Set(mv)

	return nil
}

// Exec executes query without returning rows.
func Exec(db *sql.DB, s *SQL) error {
	_, err := db.Exec(s.String())
	if err != nil {
		return NewError(opSQLDoExec, KindDatabase, err)
	}
	return nil
}

func columnName(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return ConvertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}

func setToModel(mv *reflect.Value, mt reflect.Type, indexMap *map[int]int, rVal [][]byte) error {
	// Generate reflect type and value for model.
	t := mt.Elem()
	v := reflect.Indirect(reflect.New(t))

	// Loop with number of columns in rows.
	for ri := 0; ri < len(rVal); ri++ {
		// mi is index of model field.
		mi := (*indexMap)[ri]

		valStr := string(rVal[ri])
		if valStr == "" {
			continue
		}

		// Set values to struct fields.
		if err := setField(v.Field(mi), t.Field(mi), valStr); err != nil {
			return err
		}
	}

	// Append struct to slice (or array).
	*mv = reflect.Append(*mv, v)
	return nil
}

func setField(f reflect.Value, sf reflect.StructField, v string) error {
	if !f.CanSet() {
		err := errors.New("field cannot be changes")
		return NewError(opSetField, KindBasic, err)
	}

	switch f.Kind() {
	case reflect.String:
		f.SetString(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid with value "%s"`, f.Kind(), v)
			return NewError(opSetField, KindType, err)
		}
		f.SetInt(i64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid with value "%s"`, f.Kind(), v)
			return NewError(opSetField, KindType, err)

		}
		f.SetUint(u64)
	case reflect.Float32, reflect.Float64:
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid with value "%s"`, f.Kind(), v)
			return NewError(opSetField, KindType, err)

		}
		f.SetFloat(f64)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid with value "%s"`, f.Kind(), v)
			return NewError(opSetField, KindType, err)

		}
		f.SetBool(b)
	case reflect.Struct:
		if f.Type() == reflect.TypeOf(time.Time{}) {
			layout := timeFormat(sf.Tag.Get("layout"))
			if layout == "" {
				layout = time.RFC3339
			}
			t, err := time.Parse(layout, v)
			if err != nil {
				err := fmt.Errorf(`Cannot parse %s to time.Time with format of %s`, v, layout)
				return NewError(opSetField, KindType, err)

			}
			f.Set(reflect.ValueOf(t))
		}
	}

	return nil
}

func timeFormat(layout string) string {
	switch layout {
	case "time.ANSIC":
		return time.ANSIC
	case "time.UnixDate":
		return time.UnixDate
	case "time.RubyDate":
		return time.RubyDate
	case "time.RFC822":
		return time.RFC822
	case "time.RFC822Z":
		return time.RFC822Z
	case "time.RFC850":
		return time.RFC850
	case "time.RFC1123":
		return time.RFC1123
	case "time.RFC1123Z":
		return time.RFC1123Z
	case "time.RFC3339":
		return time.RFC3339
	case "time.RFC3339Nano":
		return time.RFC3339Nano
	case "time.Kitchen":
		return time.Kitchen
	case "time.Stamp":
		return time.Stamp
	case "time.StampMilli":
		return time.StampMilli
	case "time.StampMicro":
		return time.StampMicro
	case "time.StampNano":
		return time.StampNano
	}
	return layout
}
