package database

import (
	"reflect"
	"time"
)

type sqlDriver int

// LookupDefaultType converts predefined type to database column type.
func (d sqlDriver) LookupDefaultType(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.String:
		return "VARCHAR(128)"
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return "INT"
	case reflect.Float32, reflect.Float64:
		if d == MysqlDriver {
			return "FLOAT"
		}
		if d == PsqlDriver {
			return "NUMERIC"
		}
	case reflect.Struct:
		if typ == reflect.TypeOf(time.Time{}) {
			return "DATE"
		}
	case reflect.Bool:
		return "SMALLINT"
	}
	return ""
}

// Supported SQL driver list.
const (
	PsqlDriver sqlDriver = iota
	MysqlDriver
)
