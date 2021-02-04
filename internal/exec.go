package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Op values for error handling.
const (
	opQuery           Op = "internal.Query"
	opExec            Op = "internal.Exec"
	opSetValueToField Op = "internal.setValueToField"
	opSetValueToVar   Op = "internal.setValueToVar"
)

// Query executes query and sets rows to model structure.
func Query(db *sql.DB, s *SQL, model interface{}) error {
	// Execute query.
	rows, err := db.Query(s.String())
	if err != nil {
		return NewError(opQuery, KindDatabase, err)
	}
	if rows == nil {
		return NewError(opQuery, KindDatabase, errors.New("rows is nil"))
	}
	defer rows.Close()

	// Type of model.
	mt := reflect.TypeOf(model).Elem()

	// Get columns from rows.
	rCols, err := rows.Columns()
	if err != nil {
		return NewError(opQuery, KindDatabase, err)
	}

	// Prepare pointers which is used to rows.Scan().
	rVal := make([][]byte, len(rCols))
	rValPtr := []interface{}{}
	for i := 0; i < len(rVal); i++ {
		rValPtr = append(rValPtr, &rVal[i])
	}

	switch mt.Kind() {
	case reflect.Slice, reflect.Array:
		// Generate new slice|array.
		vec := reflect.New(mt).Elem()

		// Get index map.
		indR2M := mapOfColumnsToFields(rCols, mt.Elem())

		// Loop with rows.
		for rows.Next() {
			// Scan values from rows.
			if err := rows.Scan(rValPtr...); err != nil {
				return NewError(opQuery, KindDatabase, err)
			}

			// Set values to model struct.
			v := reflect.New(mt.Elem()).Elem()
			if err := setValuesToModel(v, &indR2M, rVal); err != nil {
				return err
			}

			// Append value to slice|array.
			vec = reflect.Append(vec, v)
		}

		// Set slice|array to model.
		ref := reflect.ValueOf(model).Elem()
		ref.Set(vec)
	case reflect.Struct:
		// Generate new struct.
		v := reflect.New(mt).Elem()

		// Get index map.
		indR2M := mapOfColumnsToFields(rCols, mt)

		// Scan values from rows.
		if rows.Next() {
			if err := rows.Scan(rValPtr...); err != nil {
				return NewError(opQuery, KindDatabase, err)
			}

			// Set values to model struct.
			if err := setValuesToModel(v, &indR2M, rVal); err != nil {
				return err
			}
		}

		// Set to model.
		ref := reflect.ValueOf(model).Elem()
		ref.Set(v)
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
		// Generate new variable.
		v := reflect.New(mt).Elem()

		if rows.Next() {
			// Scan values from rows.
			if err := rows.Scan(rValPtr...); err != nil {
				return NewError(opQuery, KindDatabase, err)
			}

			// Set values to model.
			valStr := string(rVal[0])
			if err := setValueToVar(v, valStr); err != nil {
				return err
			}
		}

		ref := reflect.ValueOf(model).Elem()
		ref.Set(v)
	default:
		err := fmt.Errorf("Type %v is not supported", mt.Kind())
		return NewError(opQuery, KindType, err)
	}

	return nil
}

// Exec executes query without returning rows.
func Exec(db *sql.DB, s *SQL) error {
	_, err := db.Exec(s.String())
	if err != nil {
		return NewError(opExec, KindDatabase, err)
	}
	return nil
}

// mapOfColumnsToFields returns map to localize between column and field.
func mapOfColumnsToFields(cols []string, modelTyp reflect.Type) map[int]int {
	indR2M := make(map[int]int)
	for i, c := range cols {
		for j := 0; j < modelTyp.NumField(); j++ {
			if c != columnNameFromTag(modelTyp.Field(j)) {
				continue
			}
			indR2M[i] = j
		}
	}
	return indR2M
}

// columnNameFromTag gets column name from struct field tag.
func columnNameFromTag(sf reflect.StructField) string {
	if sf.Tag.Get("mgorm") == "" {
		return ConvertToSnakeCase(sf.Name)
	}
	return sf.Tag.Get("mgorm")
}

// setValuesToModel sets values to model fields.
func setValuesToModel(ref reflect.Value, indexMap *map[int]int, rVal [][]byte) error {
	// Generate new value from type.
	v := reflect.New(reflect.TypeOf(ref.Interface())).Elem()

	// Loop with values.
	for ri := 0; ri < len(rVal); ri++ {
		// mi is index of model field.
		mi := (*indexMap)[ri]

		// Convert value to string.
		valStr := string(rVal[ri])
		if valStr == "" {
			continue
		}

		// Set value to struct field.
		if err := setValueToField(v, mi, valStr); err != nil {
			return err
		}
	}

	// Set to model.
	ref.Set(v)
	return nil
}

// setValueToField sets string value to struct field.
func setValueToField(modelRef reflect.Value, index int, val string) error {
	// Get field from model.
	ref := modelRef.Field(index)
	if !ref.CanSet() {
		err := errors.New("Cannot set to field")
		return NewError(opSetValueToField, KindBasic, err)
	}

	switch ref.Kind() {
	case reflect.String:
		ref.SetString(val)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(ref, val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(ref, val)
	case reflect.Float32, reflect.Float64:
		return setFloat(ref, val)
	case reflect.Bool:
		return setBool(ref, val)
	case reflect.Struct:
		if ref.Type() == reflect.TypeOf(time.Time{}) {
			sf := reflect.TypeOf(modelRef.Interface()).Field(index)
			layout := TimeFormat(sf.Tag.Get("layout"))
			if layout == "" {
				layout = time.RFC3339
			}
			return setTime(ref, val, layout)
		}
	}

	err := fmt.Errorf("Type %v is not supported", ref.Kind())
	return NewError(opSetValueToField, KindType, err)
}

// setValueToVar sets string value to variable.
func setValueToVar(ref reflect.Value, val string) error {
	if !ref.CanSet() {
		err := errors.New("Cannot set to variable")
		return NewError(opSetValueToField, KindBasic, err)
	}

	switch ref.Kind() {
	case reflect.String:
		ref.SetString(val)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(ref, val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(ref, val)
	case reflect.Float32, reflect.Float64:
		return setFloat(ref, val)
	case reflect.Bool:
		return setBool(ref, val)
	case reflect.Struct:
		if ref.Type() == reflect.TypeOf(time.Time{}) {
			return setTime(ref, val, time.RFC3339)
		}
	}

	err := fmt.Errorf("Type %v is not supported", ref.Kind())
	return NewError(opSetValueToVar, KindType, err)
}

func setInt(ref reflect.Value, val string) error {
	i64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		err := fmt.Errorf(`Field type "%v" is invalid with value "%s"`, ref.Kind(), val)
		return NewError(opSetValueToField, KindType, err)
	}
	ref.SetInt(i64)
	return nil
}

func setUint(ref reflect.Value, val string) error {
	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		err := fmt.Errorf(`Field type "%v" is invalid with value "%s"`, ref.Kind(), val)
		return NewError(opSetValueToField, KindType, err)

	}
	ref.SetUint(u64)
	return nil
}

func setFloat(ref reflect.Value, val string) error {
	f64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		err := fmt.Errorf(`Field type "%v" is invalid with value "%s"`, ref.Kind(), val)
		return NewError(opSetValueToField, KindType, err)

	}
	ref.SetFloat(f64)
	return nil
}

func setBool(ref reflect.Value, val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		err := fmt.Errorf(`Field type "%v" is invalid with value "%s"`, ref.Kind(), val)
		return NewError(opSetValueToField, KindType, err)

	}
	ref.SetBool(b)
	return nil
}

func setTime(ref reflect.Value, val string, layout string) error {
	t, err := time.Parse(layout, val)
	if err != nil {
		err := fmt.Errorf(`Cannot parse %s to time.Time with format of %s`, val, layout)
		return NewError(opSetValueToField, KindType, err)

	}
	ref.Set(reflect.ValueOf(t))
	return nil
}
