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
	opSQLDoQuery internal.Op = "mgorm.SQL.doQuery"
	opSQLDoExec  internal.Op = "mgorm.SQL.doExec"
	opSetField   internal.Op = "mgorm.setField"
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

// doQuery executes query and sets rows to model structure.
func (s *SQL) doQuery(db sqlDB, model interface{}) error {
	// Execute query.
	rows, err := db.query(s.string())
	if err != nil {
		return internal.NewError(opSQLDoQuery, internal.KindDatabase, err)
	}
	if rows == nil {
		return internal.NewError(opSQLDoQuery, internal.KindDatabase, errors.New("rows is nil"))
	}

	// Model reflection.
	mt := reflect.TypeOf(model).Elem()
	mv := reflect.New(mt).Elem()

	// Model type must be slice or array.
	if mt == nil || (mt.Kind() != reflect.Slice && mt.Kind() != reflect.Array) {
		err := errors.New("model type must be slice or array")
		return internal.NewError(opSQLDoQuery, internal.KindType, err)
	}

	// Generate map to localize field index between row and model.
	rCols, err := rows.Columns()
	indR2M := make(map[int]int)
	if err != nil {
		return internal.NewError(opSQLDoQuery, internal.KindDatabase, err)
	}
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
			return internal.NewError(opSQLDoQuery, internal.KindDatabase, err)
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

// doExec executes query without returning rows.
func (s *SQL) doExec(db sqlDB) error {
	_, err := db.exec(s.string())
	if err != nil {
		return internal.NewError(opSQLDoExec, internal.KindDatabase, err)
	}
	return nil
}

func setToModel(mv *reflect.Value, mt reflect.Type, indR2M *map[int]int, rVal [][]byte) error {
	// Generate reflect type and value for model.
	t := mt.Elem()
	v := reflect.Indirect(reflect.New(t))

	// Loop with number of columns in rows.
	for ri := 0; ri < len(rVal); ri++ {
		// mi is index of model field.
		mi := (*indR2M)[ri]

		// Set values to struct fields.
		if err := setField(v.Field(mi), t.Field(mi), rVal[ri]); err != nil {
			return err
		}
	}

	// Append struct to slice (or array).
	*mv = reflect.Append(*mv, v)
	return nil
}

func columnName(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return internal.ConvertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}

func setField(f reflect.Value, sf reflect.StructField, v []byte) error {
	if !f.CanSet() {
		err := errors.New("field cannot be changes")
		return internal.NewError(opSetField, internal.KindBasic, err)
	}

	switch f.Kind() {
	case reflect.String:
		sv := string(v)
		f.SetString(sv)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src := string(v)
		i64, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid`, f.Kind())
			return internal.NewError(opSetField, internal.KindType, err)
		}
		f.SetInt(i64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src := string(v)
		u64, err := strconv.ParseUint(src, 10, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid`, f.Kind())
			return internal.NewError(opSetField, internal.KindType, err)

		}
		f.SetUint(u64)
	case reflect.Float32, reflect.Float64:
		src := string(v)
		f64, err := strconv.ParseFloat(src, 64)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid`, f.Kind())
			return internal.NewError(opSetField, internal.KindType, err)

		}
		f.SetFloat(f64)
	case reflect.Bool:
		src := string(v)
		b, err := strconv.ParseBool(src)
		if err != nil {
			err := fmt.Errorf(`field type "%v" is invalid`, f.Kind())
			return internal.NewError(opSetField, internal.KindType, err)

		}
		f.SetBool(b)
	case reflect.Struct:
		if f.Type() == reflect.TypeOf(time.Time{}) {
			src := string(v)
			layout := timeFormat(sf.Tag.Get("layout"))
			if layout == "" {
				layout = time.RFC3339
			}
			t, err := time.Parse(layout, src)
			if err != nil {
				err := fmt.Errorf(`field type "%v" is invalid`, f.Kind())
				return internal.NewError(opSetField, internal.KindType, err)

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
