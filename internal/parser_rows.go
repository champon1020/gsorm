package internal

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/champon1020/mgorm/errors"
)

// RowsParser is parser for sql.Rows.
type RowsParser struct {
	// SQL rows.
	rows *sql.Rows

	// Column names.
	Cols []string

	// Column values from sql.Rows.
	Vals [][]byte

	// Pointers of Vals.
	ValsPtr []interface{}

	// Type of model.
	Model reflect.Type

	// Correspondance between column names and struct fields.
	ColumnField map[int]int

	// Bytes parser.
	BytesParser *BytesParser

	// Error.
	Error error
}

// NewRowsParser creates RowsParser instance.
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
		rows:        rows,
		Cols:        cols,
		Vals:        vals,
		ValsPtr:     valsPtr,
		Model:       mTyp.Elem(),
		BytesParser: NewBytesParser(),
	}
	return parser, nil
}

// Next advances to next rows.
func (p *RowsParser) Next() bool {
	next := p.rows.Next()
	if err := p.rows.Scan(p.ValsPtr...); err != nil {
		p.Error = err
		return false
	}
	return next
}

// Parse converts sql.Rows to reflect.Value.
func (p *RowsParser) Parse() (*reflect.Value, error) {
	switch p.Model.Kind() {
	case reflect.Slice, reflect.Array:
		if p.Model.Elem().Kind() == reflect.Struct {
			return p.ParseStructSlice(p.Model)
		}
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
	return nil, nil
}

// ParseStructSlice converts slice or array of struct to reflect.Value.
func (p *RowsParser) ParseStructSlice(target reflect.Type) (*reflect.Value, error) {
	p.ColumnField = p.columnsAndFields(target.Elem())
	slice := reflect.New(target).Elem()

	for p.Next() {
		item, err := p.ParseStruct(target.Elem())
		if err != nil {
			return nil, err
		}
		slice = reflect.Append(slice, *item)
	}

	return &slice, nil
}

// ParseSlice converts slice or array to reflect.Value.
// If the type of elements of slice or array is struct, ParseStructSlice should be used.
func (p *RowsParser) ParseSlice(target reflect.Type) (*reflect.Value, error) {
	if len(p.Cols) != 1 {
		msg := fmt.Sprintf("Column length must be 1 but got %d", len(p.Cols))
		return nil, errors.New(msg, errors.DBColumnError)
	}

	slice := reflect.New(target).Elem()
	for p.Next() {
		item, err := p.ParseSlice(target.Elem())
		if err != nil {
			return nil, err
		}
		slice = reflect.Append(slice, *item)
	}

	return &slice, nil
}

// ParseStruct converts struct to reflect.Value.
func (p *RowsParser) ParseStruct(target reflect.Type) (*reflect.Value, error) {
	if p.ColumnField == nil {
		p.ColumnField = p.columnsAndFields(target)
	}

	item := reflect.New(target).Elem()
	for i := 0; i < len(p.Vals); i++ {
		fIdx := p.ColumnField[i]
		if !item.Field(fIdx).CanSet() {
			msg := fmt.Sprintf("Cannot set to field %d of %s", fIdx, target.String())
			return nil, errors.New(msg, errors.UnchangeableError)
		}

		opt := BytesParserOption{}
		if item.Field(fIdx).Type() == reflect.TypeOf(time.Time{}) {
			tag := ExtractTag(item.Type().Field(fIdx))
			if tag.Lookup("layout") {
				opt.TimeLayout = tag.Layout
			}
		}

		v, err := p.BytesParser.Parse(p.Vals[i], item.Field(fIdx).Type(), opt)
		if err != nil {
			return nil, err
		}

		item.Field(fIdx).Set(*v)
	}

	return &item, nil
}

// ParseVar converts variable to reflect.Value.
func (p *RowsParser) ParseVar(target reflect.Type) (*reflect.Value, error) {
	if len(p.Cols) != 1 {
		msg := fmt.Sprintf("Column length must be 1 but got %d", len(p.Cols))
		return nil, errors.New(msg, errors.DBColumnError)
	}

	item, err := p.BytesParser.Parse(p.Vals[0], target)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (p *RowsParser) columnsAndFields(target reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, col := range p.Cols {
		for j := 0; j < target.NumField(); j++ {
			c := ExtractTag(target.Field(j)).Column
			if c == "" {
				c = SnakeCase(target.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}
