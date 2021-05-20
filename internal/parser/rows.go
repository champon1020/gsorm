package parser

import (
	"database/sql"
	"reflect"
	"strconv"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

// Rows is interface for *sql.Rows.
type Rows interface {
	//	ColumnTypes() ([]*sql.ColumnType, error)
	Next() bool
	Scan(...interface{}) error
}

// ColumnType is interface for *sql.ColumnType.
type ColumnType interface {
	Name() string
	ScanType() reflect.Type
}

// RowsParser is parser for sql.Rows.
type RowsParser struct {
	// SQL rows.
	Rows Rows

	// Number of columns.
	NumOfColumns int

	// Column types
	ColumnTypes []ColumnType

	// Pointers of the scanning values.
	ItemPtr []interface{}

	// Type of the model.
	ModelType reflect.Type

	// Error.
	Error error
}

// NewRowsParser creates RowsParser instance.
func NewRowsParser(r Rows, ct []ColumnType, m interface{}) (*RowsParser, error) {
	mt := reflect.TypeOf(m)
	if mt.Kind() != reflect.Ptr {
		err := failure.New(errInvalidValue, failure.Message("model must be a pointer"))
		return nil, err
	}
	mt = mt.Elem()

	ptrs := make([]interface{}, len(ct))

	p := &RowsParser{
		Rows:         r,
		NumOfColumns: len(ct),
		ColumnTypes:  ct,
		ItemPtr:      ptrs,
		ModelType:    mt,
	}

	return p, nil
}

// Next advances to next rows.
func (p *RowsParser) Next() bool {
	if !p.Rows.Next() {
		return false
	}
	if err := p.Rows.Scan(p.ItemPtr...); err != nil {
		p.Error = err
		return false
	}
	return true
}

// Parse converts sql.Rows to reflect.Value.
func (p *RowsParser) Parse() (*reflect.Value, error) {
	switch p.ModelType.Kind() {
	case reflect.Slice,
		reflect.Array:
		// If the type of item is struct.
		if p.ModelType.Elem().Kind() == reflect.Struct {
			return p.ParseStructSlice()
		}

		// If the type of item is map.
		if p.ModelType.Elem().Kind() == reflect.Map {
			return p.ParseMapSlice()
		}

		// If the type of item is predeclared types.
		return p.ParseSlice()
	case reflect.Struct:
		return p.ParseStruct()
	case reflect.Map:
		return p.ParseMap()
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
		return p.ParseVar()
	}

	err := failure.New(errInvalidType,
		failure.Context{"type": p.ModelType.Kind().String()},
		failure.Message("invalid type for internal.InsertModelParser.Parse"))
	return nil, err
}

// ParseMapSlice converts slice or array of map to reflect.Value.
func (p *RowsParser) ParseMapSlice() (*reflect.Value, error) {
	item := make([]reflect.Value, p.NumOfColumns)
	for i := 0; i < p.NumOfColumns; i++ {
		item[i] = generateValue(p.ColumnTypes[i].ScanType())
		p.ItemPtr[i] = item[i].Addr().Interface()
	}

	sl := reflect.New(p.ModelType).Elem()
	for p.Next() {
		mp := reflect.MakeMap(p.ModelType.Elem())
		for i := 0; i < p.NumOfColumns; i++ {
			k := reflect.ValueOf(p.ColumnTypes[i].Name())
			mp.SetMapIndex(k, item[i])
		}
		sl = reflect.Append(sl, mp)
	}

	if p.Error != nil {
		return nil, p.Error
	}

	return &sl, nil
}

// ParseStructSlice converts the slice or array of struct to reflect.Value.
// dest is the destination type. In this case, underlying type of dest is struct.
func (p *RowsParser) ParseStructSlice() (*reflect.Value, error) {
	item := reflect.New(p.ModelType.Elem()).Elem()
	cf := p.columnsAndFields(p.ModelType.Elem())
	for i := 0; i < p.NumOfColumns; i++ {
		p.ItemPtr[i] = item.Field(cf[i]).Addr().Interface()
	}

	sl := reflect.New(p.ModelType).Elem()
	for p.Next() {
		sl = reflect.Append(sl, item)
	}

	if p.Error != nil {
		return nil, p.Error
	}

	return &sl, nil
}

// ParseSlice converts slice or array to reflect.Value.
// If the type of elements of slice or array is struct or map, ParseStructSlice or ParseMapSlice should be used.
func (p *RowsParser) ParseSlice() (*reflect.Value, error) {
	if p.NumOfColumns != 1 {
		err := failure.New(errInvalidSyntax,
			failure.Context{"numColumns": strconv.Itoa(p.NumOfColumns)},
			failure.Message("invalid number of columns"))
		return nil, err
	}

	item := reflect.New(p.ModelType.Elem()).Elem()
	p.ItemPtr[0] = item.Addr().Interface()

	sl := reflect.New(p.ModelType).Elem()
	for p.Next() {
		sl = reflect.Append(sl, item)
	}

	if p.Error != nil {
		return nil, p.Error
	}

	return &sl, nil
}

// ParseMap converts map to reflect.Value.
func (p *RowsParser) ParseMap() (*reflect.Value, error) {
	item := make([]reflect.Value, p.NumOfColumns)
	for i := 0; i < p.NumOfColumns; i++ {
		item[i] = generateValue(p.ColumnTypes[i].ScanType())
		p.ItemPtr[i] = item[i].Addr().Interface()
	}
	p.Next()
	mp := reflect.MakeMap(p.ModelType)
	for i := 0; i < p.NumOfColumns; i++ {
		k := reflect.ValueOf(p.ColumnTypes[i].Name())
		mp.SetMapIndex(k, item[i])
	}
	return &mp, nil
}

// ParseStruct converts struct to reflect.Value.
func (p *RowsParser) ParseStruct() (*reflect.Value, error) {
	item := reflect.New(p.ModelType).Elem()
	cf := p.columnsAndFields(p.ModelType)
	for i := 0; i < p.NumOfColumns; i++ {
		p.ItemPtr[i] = item.Field(cf[i]).Addr().Interface()
	}
	p.Next()
	return &item, nil
}

// ParseVar converts variable to reflect.Value.
func (p *RowsParser) ParseVar() (*reflect.Value, error) {
	if p.NumOfColumns != 1 {
		err := failure.New(errInvalidSyntax,
			failure.Context{"numColumns": strconv.Itoa(p.NumOfColumns)},
			failure.Message("invalid number of columns"))
		return nil, err
	}

	item := reflect.New(p.ModelType).Elem()
	p.ItemPtr[0] = item.Addr().Interface()
	p.Next()
	return &item, nil
}

func (p *RowsParser) columnsAndFields(dest reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, ct := range p.ColumnTypes {
		for j := 0; j < dest.NumField(); j++ {
			c := internal.ExtractTag(dest.Field(j)).Column
			if c == "" {
				c = internal.SnakeCase(dest.Field(j).Name)
			}
			if ct.Name() != c {
				continue
			}
			cf[i] = j
		}
	}
	return cf
}

// Types.
var (
	Int         = reflect.TypeOf(int(0))
	Int8        = reflect.TypeOf(int8(0))
	Int16       = reflect.TypeOf(int16(0))
	Int32       = reflect.TypeOf(int32(0))
	Int64       = reflect.TypeOf(int64(0))
	Uint8       = reflect.TypeOf(uint8(0))
	Uint16      = reflect.TypeOf(uint16(0))
	Uint32      = reflect.TypeOf(uint32(0))
	Uint64      = reflect.TypeOf(uint64(0))
	Float32     = reflect.TypeOf(float32(0))
	Float64     = reflect.TypeOf(float64(0))
	Bool        = reflect.TypeOf(false)
	String      = reflect.TypeOf("")
	NullInt32   = reflect.TypeOf(sql.NullInt32{})
	NullInt64   = reflect.TypeOf(sql.NullInt64{})
	NullFloat64 = reflect.TypeOf(sql.NullFloat64{})
	NullBool    = reflect.TypeOf(sql.NullBool{})
	NullString  = reflect.TypeOf(sql.NullString{})
	NullTime    = reflect.TypeOf(sql.NullTime{})
	RawBytes    = reflect.TypeOf(sql.RawBytes{})
)

func generateValue(t reflect.Type) reflect.Value {
	switch t {
	case Int8, Int16, Int32, Int64, Uint8, Uint16, Uint32, Uint64:
		return reflect.New(Int).Elem()
	case RawBytes:
		return reflect.New(String).Elem()
	}

	// Bool, String, Float32, Float64
	// NullInt32, NullInt64, NullBool, NullFloat64, NullString, NullTime
	return reflect.New(t).Elem()
}
