package internal_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

type Rows struct {
	Cols  []string
	Vals  [][]string
	count int
}

func (r *Rows) Columns() ([]string, error) {
	return r.Cols, nil
}

func (r *Rows) Next() bool {
	r.count++
	return r.count-1 < len(r.Vals)
}

func (r *Rows) Scan(dest ...interface{}) error {
	for i, d := range dest {
		b := d.(*[]byte)
		*b = []byte(r.Vals[r.count-1][i])
	}
	return nil
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

func TestRowsParser_ParseMapSlice(t *testing.T) {
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
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
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseMapSlice(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseStructSlice(t *testing.T) {
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
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
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseStructSlice(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseSlice(t *testing.T) {
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int"},
			[][]string{{"9223372036854775807"}, {"-9223372036854775808"}},
			&[]int{},
			[]int{9223372036854775807, -9223372036854775808},
		},
	}

	for _, testCase := range testCases {
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		v, err := p.ParseSlice(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseMap(t *testing.T) {
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
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
			[]string{"float32", "float64"},
			[][]string{{"3.14159265358979", "3.141592653589793238462643383279"}},
			&map[string]interface{}{},
			map[string]interface{}{
				"float32": 3.14159265358979,
				"float64": 3.141592653589793238462643383279,
			},
		},
		{
			[]string{"bool", "time"},
			[][]string{{"true", "2021-01-02T03:04:05Z"}},
			&map[string]interface{}{},
			map[string]interface{}{
				"bool": true,
				"time": time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseMap(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}

func TestRowsParser_ParseStruct(t *testing.T) {
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
			[][]string{{"-9223372036854775808", "127", "32767", "2147483647", "9223372036854775807"}},
			&IntModel{},
			IntModel{I: -9223372036854775808, I8: 127, I16: 32767, I32: 2147483647, I64: 9223372036854775807},
		},
		{
			[]string{"uint", "uint8", "uint16", "uint32", "uint64"},
			[][]string{{"1", "255", "65535", "4294967295", "18446744073709551615"}},
			&UintModel{},
			UintModel{U: 1, U8: 255, U16: 65535, U32: 4294967295, U64: 18446744073709551615},
		},
		{
			[]string{"float32", "float64"},
			[][]string{{"3.14159265358979", "3.141592653589793238462643383279"}},
			&FloatModel{},
			FloatModel{F32: 3.14159265358979, F64: 3.141592653589793238462643383279},
		},
		{
			[]string{"bool", "time", "time_ansic", "time_format"},
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
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseStruct(p.Model)
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
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int"},
			[][]string{{"-9223372036854775808"}},
			&iModel.I,
			-1 << 63,
		},
		{
			[]string{"int8"},
			[][]string{{"127"}},
			&iModel.I8,
			int8(1<<7 - 1),
		},
		{
			[]string{"int16"},
			[][]string{{"32767"}},
			&iModel.I16,
			int16(1<<15 - 1),
		},
		{
			[]string{"int32"},
			[][]string{{"2147483647"}},
			&iModel.I32,
			int32(1<<31 - 1),
		},
		{
			[]string{"int64"},
			[][]string{{"9223372036854775807"}},
			&iModel.I64,
			int64(1<<63 - 1),
		},
		{
			[]string{"uint"},
			[][]string{{"1"}},
			&uModel.U,
			uint(1),
		},
		{
			[]string{"uint8"},
			[][]string{{"255"}},
			&uModel.U8,
			uint8(1<<8 - 1),
		},
		{
			[]string{"uint16"},
			[][]string{{"65535"}},
			&uModel.U16,
			uint16(1<<16 - 1),
		},
		{
			[]string{"uint32"},
			[][]string{{"4294967295"}},
			&uModel.U32,
			uint32(1<<32 - 1),
		},
		{
			[]string{"uint64"},
			[][]string{{"18446744073709551615"}},
			&uModel.U64,
			uint64(1<<64 - 1),
		},
		{
			[]string{"float32"},
			[][]string{{"3.14159265358979"}},
			&fModel.F32,
			float32(3.14159265358979),
		},
		{
			[]string{"float64"},
			[][]string{{"3.141592653589793238462643383279"}},
			&fModel.F64,
			float64(3.141592653589793238462643383279),
		},
		{
			[]string{"bool"},
			[][]string{{"true"}},
			&othersModel.B,
			true,
		},
		{
			[]string{"time"},
			[][]string{{"2021-01-02T03:04:05Z"}},
			&othersModel.Time,
			time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, testCase := range testCases {
		p, err := internal.NewRowsParser(&Rows{Cols: testCase.RowsCols, Vals: testCase.RowsVals}, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}

		p.Next()
		v, err := p.ParseVar(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}
