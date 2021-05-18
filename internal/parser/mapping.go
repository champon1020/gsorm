package parser

import (
	"database/sql"
	"reflect"
)

// MapRowsToModel executes query and sets rows to model structure.
func MapRowsToModel(rows *sql.Rows, model interface{}) error {
	p, err := NewRowsParser(rows, model)
	if err != nil {
		return err
	}

	v, err := p.Parse()
	if err != nil {
		return err
	}

	ref := reflect.ValueOf(model).Elem()
	ref.Set(*v)
	return nil
}