package parser

import (
	"reflect"
	"strconv"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

// Rows is interface for *sql.Rows.
type Rows interface {
	Columns() ([]string, error)
	Next() bool
	Scan(...interface{}) error
}

// RowsParser is parser for sql.Rows.
type RowsParser struct {
	// SQL rows.
	Rows Rows

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
func NewRowsParser(rows Rows, model interface{}) (*RowsParser, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, failure.Wrap(err)
	}

	vals := make([][]byte, len(cols))
	valsPtr := make([]interface{}, 0, len(cols))
	for i := 0; i < len(vals); i++ {
		valsPtr = append(valsPtr, &vals[i])
	}

	mTyp := reflect.TypeOf(model)
	if mTyp.Kind() != reflect.Ptr {
		err := failure.New(errInvalidValue, failure.Message("model must be a pointer"))
		return nil, err
	}

	parser := &RowsParser{
		Rows:        rows,
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
	if !p.Rows.Next() {
		return false
	}
	if err := p.Rows.Scan(p.ValsPtr...); err != nil {
		p.Error = err
		return false
	}
	return true
}

// Parse converts sql.Rows to reflect.Value.
func (p *RowsParser) Parse() (*reflect.Value, error) {
	switch p.Model.Kind() {
	case reflect.Slice,
		reflect.Array:
		if p.Model.Elem().Kind() == reflect.Struct {
			return p.ParseStructSlice(p.Model)
		}
		if p.Model.Elem().Kind() == reflect.Map {
			return p.ParseMapSlice(p.Model)
		}
		return p.ParseSlice(p.Model)
	case reflect.Struct:
		p.Next()
		return p.ParseStruct(p.Model)
	case reflect.Map:
		p.Next()
		return p.ParseMap(p.Model)
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
		p.Next()
		return p.ParseVar(p.Model)
	}

	err := failure.New(errInvalidType,
		failure.Context{"type": p.Model.Kind().String()},
		failure.Message("invalid type for internal.InsertModelParser.Parse"))
	return nil, err
}

// ParseMapSlice converts slice or array of map to reflect.Value.
func (p *RowsParser) ParseMapSlice(target reflect.Type) (*reflect.Value, error) {
	slice := reflect.New(target).Elem()

	for p.Next() {
		item, err := p.ParseMap(target.Elem())
		if err != nil {
			return nil, err
		}
		slice = reflect.Append(slice, *item)
	}

	if p.Error != nil {
		return nil, p.Error
	}

	return &slice, nil
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

	if p.Error != nil {
		return nil, p.Error
	}

	return &slice, nil
}

// ParseSlice converts slice or array to reflect.Value.
// If the type of elements of slice or array is struct or map, ParseStructSlice or ParseMapSlice should be used.
func (p *RowsParser) ParseSlice(target reflect.Type) (*reflect.Value, error) {
	if len(p.Cols) != 1 {
		err := failure.New(errInvalidSyntax,
			failure.Context{"numColumns": strconv.Itoa(len(p.Cols))},
			failure.Message("invalid number of columns"))
		return nil, err
	}

	slice := reflect.New(target).Elem()
	for p.Next() {
		item, err := p.ParseVar(target.Elem())
		if err != nil {
			return nil, err
		}
		slice = reflect.Append(slice, *item)
	}

	if p.Error != nil {
		return nil, p.Error
	}

	return &slice, nil
}

// ParseMap converts map to reflect.Value.
func (p *RowsParser) ParseMap(target reflect.Type) (*reflect.Value, error) {
	item := reflect.MakeMap(target)
	for i := 0; i < len(p.Vals); i++ {
		key := reflect.ValueOf(p.Cols[i])
		val := p.BytesParser.ParseAuto(p.Vals[i])
		item.SetMapIndex(key, *val)
	}
	return &item, nil
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
			err := failure.New(errInvalidSyntax,
				failure.Context{"model": target.String(), "fieldIndex": strconv.Itoa(fIdx)},
				failure.Message("cannot change the value of the field"))
			return nil, err
		}

		opt := BytesParserOption{}
		if item.Field(fIdx).Type() == reflect.TypeOf(time.Time{}) {
			tag := internal.ExtractTag(item.Type().Field(fIdx))
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
		err := failure.New(errInvalidSyntax,
			failure.Context{"numColumns": strconv.Itoa(len(p.Cols))},
			failure.Message("invalid number of columns"))
		return nil, err
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
			c := internal.ExtractTag(target.Field(j)).Column
			if c == "" {
				c = internal.SnakeCase(target.Field(j).Name)
			}
			if col != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}