package gsorm

import (
	"database/sql"
	"reflect"

	"github.com/champon1020/gsorm/internal"
	"golang.org/x/xerrors"
)

// rowsParser is parser for sql.Rows.
type rowsParser struct {
	// SQL rows.
	rows irows

	// Number of columns.
	numOfColumns int

	// Column types
	columnTypes []icolumnType

	// Pointers of the scanning values.
	itemPtr []interface{}

	// Type of the model.
	modelType reflect.Type

	// Error.
	err error
}

// newRowsParser creates rowsParser instance.
func newRowsParser(r irows, m interface{}) (*rowsParser, error) {
	ct, err := r.ColumnTypes()
	if err != nil {
		return nil, err
	}

	cti := make([]icolumnType, len(ct))
	for i := 0; i < len(ct); i++ {
		cti[i] = ct[i]
	}

	mt := reflect.TypeOf(m)
	if mt.Kind() != reflect.Ptr {
		return nil, xerrors.New("model must be a pointer")
	}
	mt = mt.Elem()

	ptrs := make([]interface{}, len(ct))

	p := &rowsParser{
		rows:         r,
		numOfColumns: len(cti),
		columnTypes:  cti,
		itemPtr:      ptrs,
		modelType:    mt,
	}

	return p, nil
}

// Next advances to next rows.
func (p *rowsParser) Next() bool {
	if !p.rows.Next() {
		return false
	}
	if err := p.rows.Scan(p.itemPtr...); err != nil {
		p.err = err
		return false
	}
	return true
}

// Parse converts sql.Rows to reflect.Value.
func (p *rowsParser) Parse() (*reflect.Value, error) {
	switch p.modelType.Kind() {
	case reflect.Slice,
		reflect.Array:
		// If the type of item is struct.
		if p.modelType.Elem().Kind() == reflect.Struct {
			return p.ParseStructSlice()
		}

		// If the type of item is map.
		if p.modelType.Elem().Kind() == reflect.Map {
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

	return nil, xerrors.Errorf("%s is invalid type for rowsParser.Parse", p.modelType.Kind().String())
}

// ParseMapSlice converts slice or array of map to reflect.Value.
func (p *rowsParser) ParseMapSlice() (*reflect.Value, error) {
	item := make([]reflect.Value, p.numOfColumns)
	for i := 0; i < p.numOfColumns; i++ {
		item[i] = generateValue(p.columnTypes[i].ScanType())
		p.itemPtr[i] = item[i].Addr().Interface()
	}

	sl := reflect.New(p.modelType).Elem()
	for p.Next() {
		mp := reflect.MakeMap(p.modelType.Elem())
		for i := 0; i < p.numOfColumns; i++ {
			k := reflect.ValueOf(p.columnTypes[i].Name())
			mp.SetMapIndex(k, item[i])
		}
		sl = reflect.Append(sl, mp)
	}

	if p.err != nil {
		return nil, p.err
	}

	return &sl, nil
}

// ParseStructSlice converts the slice or array of struct to reflect.Value.
// dest is the destination type. In this case, underlying type of dest is struct.
func (p *rowsParser) ParseStructSlice() (*reflect.Value, error) {
	item := reflect.New(p.modelType.Elem()).Elem()
	cf := p.columnsAndFields(p.modelType.Elem())
	for i := 0; i < p.numOfColumns; i++ {
		p.itemPtr[i] = item.Field(cf[i]).Addr().Interface()
	}

	sl := reflect.New(p.modelType).Elem()
	for p.Next() {
		sl = reflect.Append(sl, item)
	}

	if p.err != nil {
		return nil, p.err
	}

	return &sl, nil
}

// ParseSlice converts slice or array to reflect.Value.
// If the type of elements of slice or array is struct or map, ParseStructSlice or ParseMapSlice should be used.
func (p *rowsParser) ParseSlice() (*reflect.Value, error) {
	if p.numOfColumns != 1 {
		return nil, xerrors.Errorf("number of columns must be 1, not %d", p.numOfColumns)
	}

	item := reflect.New(p.modelType.Elem()).Elem()
	p.itemPtr[0] = item.Addr().Interface()

	sl := reflect.New(p.modelType).Elem()
	for p.Next() {
		sl = reflect.Append(sl, item)
	}

	if p.err != nil {
		return nil, p.err
	}

	return &sl, nil
}

// ParseMap converts map to reflect.Value.
func (p *rowsParser) ParseMap() (*reflect.Value, error) {
	item := make([]reflect.Value, p.numOfColumns)
	for i := 0; i < p.numOfColumns; i++ {
		item[i] = generateValue(p.columnTypes[i].ScanType())
		p.itemPtr[i] = item[i].Addr().Interface()
	}
	p.Next()
	mp := reflect.MakeMap(p.modelType)
	for i := 0; i < p.numOfColumns; i++ {
		k := reflect.ValueOf(p.columnTypes[i].Name())
		mp.SetMapIndex(k, item[i])
	}
	return &mp, nil
}

// ParseStruct converts struct to reflect.Value.
func (p *rowsParser) ParseStruct() (*reflect.Value, error) {
	item := reflect.New(p.modelType).Elem()
	cf := p.columnsAndFields(p.modelType)
	for i := 0; i < p.numOfColumns; i++ {
		p.itemPtr[i] = item.Field(cf[i]).Addr().Interface()
	}
	p.Next()
	return &item, nil
}

// ParseVar converts variable to reflect.Value.
func (p *rowsParser) ParseVar() (*reflect.Value, error) {
	if p.numOfColumns != 1 {
		return nil, xerrors.Errorf("number of columns must be 1, not %d", p.numOfColumns)
	}

	item := reflect.New(p.modelType).Elem()
	p.itemPtr[0] = item.Addr().Interface()
	p.Next()
	return &item, nil
}

func (p *rowsParser) columnsAndFields(dest reflect.Type) map[int]int {
	cf := make(map[int]int)
	for i, ct := range p.columnTypes {
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
