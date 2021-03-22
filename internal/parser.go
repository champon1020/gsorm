package internal

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
)

type RowsParser struct {
	rows    *sql.Rows
	Cols    []string
	Vals    [][]byte
	ValsPtr []interface{}
	Model   reflect.Type

	Error error
}

func NewRowsParser(rows *sql.Rows, model interface{}) (*RowsParser, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.New(err.Error(), errors.DBColumnError)
	}

	vals := make([][]byte, 0, len(cols))
	valsPtr := make([]interface{}, 0, len(cols))
	for i := 0; i < len(vals); i++ {
		valsPtr = append(valsPtr, &vals[i])
	}

	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		return nil, errors.New(err.Error(), errors.InvalidTypeError)
	}

	parser := &RowsParser{
		rows:    rows,
		Cols:    cols,
		Vals:    vals,
		ValsPtr: valsPtr,
		Model:   mTyp.Elem(),
	}
	return parser, nil
}

func (p *RowsParser) Next() bool {
	next := p.rows.Next()
	if err := p.rows.Scan(p.ValsPtr...); err != nil {
		p.Error = err
		return false
	}
	return next
}

func (p *RowsParser) Parse() reflect.Value {
	switch p.Model.Kind() {
	case reflect.Slice, reflect.Array:
	case reflect.Struct:
	case reflect.Map:
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
	}
}

func (p *RowsParser) ParseStructSlice() (*reflect.Value, error) {
	cf := p.columnsAndFields()
	slice := reflect.New(p.Model).Elem()
	for p.Next() {
		item := reflect.New(p.Model.Elem()).Elem()
		for i := 0; i < len(p.Vals); i++ {
			// Index of field.
			fIdx := cf[i]
			if !item.Field(fIdx).CanSet() {
				msg := fmt.Sprintf("Cannot set to field %d of %s", fIdx, p.Model.String())
				return nil, errors.New(msg, errors.UnchangeableError)
			}

			// Value from row.
			val := Str(p.Vals[i])
			if val.Empty() {
				continue
			}

			// Kind of field.
			fTyp := item.Field(fIdx).Kind()
		}
	}

	if p.Error != nil {
		return nil, p.Error
	}
}

func (p *RowsParser) columnsAndFields() map[int]int {
	cf := make(map[int]int)
	for i, col := range p.Cols {
		for j := 0; j < p.Model.NumField(); j++ {
			c := ExtractTag(p.Model.Field(j)).Column
			if c == "" {
				c = SnakeCase(p.Model.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}
