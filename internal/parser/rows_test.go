package parser_test

import (
	"reflect"
	"time"
)

type Rows struct {
	values [][]string
	count  int
}

func (r *Rows) Next() bool {
	r.count++
	return r.count-1 < len(r.values)
}

func (r *Rows) Scan(dest ...interface{}) error {
	for i, d := range dest {
		b := d.(*[]byte)
		*b = []byte(r.values[r.count-1][i])
	}
	return nil
}

type ColumnType struct {
	name     string
	scanType reflect.Type
}

func (c *ColumnType) Name() string {
	return c.name
}

func (c *ColumnType) ScanType() reflect.Type {
	return c.scanType
}

type IntModel struct {
	I8  int8  `mgorm:"int8"`
	I16 int16 `mgorm:"int16"`
	I32 int32 `mgorm:"int32"`
	I64 int64 `mgorm:"int64"`
	I   int   `mgorm:"int"`
}

type UintModel struct {
	U8  uint8  `mgorm:"uint8"`
	U16 uint16 `mgorm:"uint16"`
	U32 uint32 `mgorm:"uint32"`
	U64 uint64 `mgorm:"uint64"`
	U   uint   `mgorm:"uint"`
}

type FloatModel struct {
	F32 float32 `mgorm:"float32"`
	F64 float64 `mgorm:"float64"`
}

type OtherTypesModel struct {
	B          bool      `mgorm:"bool"`
	Time       time.Time `mgorm:"time"`
	TimeANSIC  time.Time `mgorm:"time_ansic,layout=time.ANSIC"`
	TimeFormat time.Time `mgorm:"time_format,layout=2006-01-02"`
}

/*
func TestRowsParser_ParseMapSlice(t *testing.T) {
	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "int", scanType: parser.NullInt64},
				&ColumnType{name: "int8", scanType: parser.Int8},
				&ColumnType{name: "int16", scanType: parser.Int16},
				&ColumnType{name: "int32", scanType: parser.Int32},
				&ColumnType{name: "int64", scanType: parser.Int64},
			},
			[][]string{
				{"1", "127", "32767", "2147483647", "9223372036854775807"},
				{"-1", "-128", "-32768", "-2147483648", "-9223372036854775808"},
			},
			&[]map[string]interface{}{},
			[]map[string]interface{}{
				{"int": 1, "int8": 127, "int16": 32767, "int32": 2147483647, "int64": 9223372036854775807},
				{"int": -1, "int8": -128, "int16": -32768, "int32": -2147483648, "int64": -9223372036854775808},
			},
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseMapSlice()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseStructSlice(t *testing.T) {
	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "int", scanType: parser.NullInt64},
				&ColumnType{name: "int8", scanType: parser.Int8},
				&ColumnType{name: "int16", scanType: parser.Int16},
				&ColumnType{name: "int32", scanType: parser.Int32},
				&ColumnType{name: "int64", scanType: parser.Int64},
			},
			[][]string{
				{"1", "127", "32767", "2147483647", "9223372036854775807"},
				{"-1", "-128", "-32768", "-2147483648", "-9223372036854775808"},
			},
			&[]IntModel{},
			[]IntModel{
				{I: 1, I8: 127, I16: 32767, I32: 2147483647, I64: 9223372036854775807},
				{I: -1, I8: -128, I16: -32768, I32: -2147483648, I64: -9223372036854775808},
			},
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseStructSlice()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseSlice(t *testing.T) {
	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "int", scanType: parser.Int64},
			},
			[][]string{{"9223372036854775807"}, {"-9223372036854775808"}},
			&[]int{},
			[]int{9223372036854775807, -9223372036854775808},
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseSlice()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseMap(t *testing.T) {
	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "int", scanType: parser.NullInt64},
				&ColumnType{name: "int8", scanType: parser.Int8},
				&ColumnType{name: "int16", scanType: parser.Int16},
				&ColumnType{name: "int32", scanType: parser.Int32},
				&ColumnType{name: "int64", scanType: parser.Int64},
			},
			[][]string{{"-9223372036854775808", "127", "32767", "2147483647", "9223372036854775807"}},
			&map[string]interface{}{},
			map[string]interface{}{
				"int":   -9223372036854775808,
				"int8":  127,
				"int16": 32767,
				"int32": 2147483647,
				"int64": 9223372036854775807,
			},
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "float64", scanType: parser.Float64},
			},
			[][]string{{"3.141592653589793238462643383279"}},
			&map[string]interface{}{},
			map[string]interface{}{
				"float64": 3.141592653589793,
			},
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "bool", scanType: parser.Bool},
				&ColumnType{name: "float64", scanType: parser.NullTime},
			},
			[][]string{{"true", "2021-01-02T03:04:05Z"}},
			&map[string]interface{}{},
			map[string]interface{}{
				"bool": true,
				"time": time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseMap()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseStruct(t *testing.T) {
	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "int", scanType: parser.NullInt64},
				&ColumnType{name: "int8", scanType: parser.Int8},
				&ColumnType{name: "int16", scanType: parser.Int16},
				&ColumnType{name: "int32", scanType: parser.Int32},
				&ColumnType{name: "int64", scanType: parser.Int64},
			},
			[][]string{{"-9223372036854775808", "127", "32767", "2147483647", "9223372036854775807"}},
			&IntModel{},
			IntModel{I: -9223372036854775808, I8: 127, I16: 32767, I32: 2147483647, I64: 9223372036854775807},
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "float32", scanType: parser.Float64},
				&ColumnType{name: "float64", scanType: parser.Float64},
			},
			[][]string{{"3.141592653589793238462643383279", "3.141592653589793238462643383279"}},
			&FloatModel{},
			FloatModel{F32: 3.1415927, F64: 3.141592653589793},
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "bool", scanType: parser.Bool},
				&ColumnType{name: "time", scanType: parser.NullTime},
				&ColumnType{name: "time_ansic", scanType: parser.NullTime},
				&ColumnType{name: "time_format", scanType: parser.NullTime},
			},
			[][]string{{"true", "2021-01-02T03:04:05Z", "Wed Mar 25 22:13:30 2021", "2021-04-01"}},
			&OtherTypesModel{},
			OtherTypesModel{
				B:          true,
				Time:       time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
				TimeANSIC:  time.Date(2021, time.March, 25, 22, 13, 30, 0, time.UTC),
				TimeFormat: time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseStruct()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseVar(t *testing.T) {
	iModel := IntModel{}
	uModel := UintModel{}
	fModel := FloatModel{}
	othersModel := OtherTypesModel{}

	testCases := []struct {
		ColumnTypes []parser.ColumnType
		RowsValues  [][]string
		Model       interface{}
		Expected    interface{}
	}{
		{
			[]parser.ColumnType{
				&ColumnType{name: "nullint", scanType: parser.NullInt64},
			},
			[][]string{{"-9223372036854775808"}},
			&iModel.I,
			-1 << 63,
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "int8", scanType: parser.Int8},
			},
			[][]string{{"127"}},
			&iModel.I8,
			int8(1<<7 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "int16", scanType: parser.Int16},
			},
			[][]string{{"32767"}},
			&iModel.I16,
			int16(1<<15 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "int32", scanType: parser.Int32},
			},
			[][]string{{"2147483647"}},
			&iModel.I32,
			int32(1<<31 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "int64", scanType: parser.Int64},
			},
			[][]string{{"9223372036854775807"}},
			&iModel.I64,
			int64(1<<63 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "uint8", scanType: parser.Uint8},
			},
			[][]string{{"255"}},
			&uModel.U8,
			uint8(1<<8 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "uint16", scanType: parser.Uint16},
			},
			[][]string{{"65535"}},
			&uModel.U16,
			uint16(1<<16 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "uint32", scanType: parser.Uint32},
			},
			[][]string{{"4294967295"}},
			&uModel.U32,
			uint32(1<<32 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "uint64", scanType: parser.Uint64},
			},
			[][]string{{"18446744073709551615"}},
			&uModel.U64,
			uint64(1<<64 - 1),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "float32", scanType: parser.Float32},
			},
			[][]string{{"3.141592653589793238462643383279"}},
			&fModel.F32,
			float32(3.1415927),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "float64", scanType: parser.Float64},
			},
			[][]string{{"3.141592653589793238462643383279"}},
			&fModel.F64,
			float64(3.141592653589793),
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "bool", scanType: parser.Bool},
			},
			[][]string{{"true"}},
			&othersModel.B,
			true,
		},
		{
			[]parser.ColumnType{
				&ColumnType{name: "time", scanType: parser.NullTime},
			},
			[][]string{{"2021-01-02T03:04:05Z"}},
			&othersModel.Time,
			time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewRowsParser(
			&Rows{values: testCase.RowsValues},
			testCase.ColumnTypes,
			testCase.Model,
		)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseVar()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}
*/
