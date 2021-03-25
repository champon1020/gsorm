package internal_test

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

type RowsParser struct {
	internal.RowsParser
	RowsVals [][]string

	count int
}

func (p *RowsParser) Next() bool {
	p.Vals = make([][]byte, 0, len(p.Cols))
	for _, v := range p.RowsVals[p.count] {
		p.Vals = append(p.Vals, []byte(v))
	}
	p.count++
	return p.count < len(p.RowsVals)
}

func NewRowsParser(cols []string, vals [][]string, model interface{}) *RowsParser {
	p := &RowsParser{RowsVals: vals}
	p.Cols = cols
	p.Model = reflect.TypeOf(model)
	p.BytesParser = internal.NewBytesParser()
	return p
}

func TestRowsParser_ParseStruct(t *testing.T) {
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

	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
			[][]string{{"1", "2", "3", "4", "5"}},
			IntModel{},
			IntModel{I: 1, I8: 2, I16: 3, I32: 4, I64: 5},
		},
		{
			[]string{"uint", "uint8", "uint16", "uint32", "uint64"},
			[][]string{{"1", "2", "3", "4", "5"}},
			UintModel{},
			UintModel{U: 1, U8: 2, U16: 3, U32: 4, U64: 5},
		},
		{
			[]string{"float32", "float64"},
			[][]string{{"1.1", "2.2"}},
			FloatModel{},
			FloatModel{F32: 1.1, F64: 2.2},
		},
		{
			[]string{"bool", "time", "time_ansic", "time_format"},
			[][]string{{"true", "2021-01-02T03:04:05Z", "Wed Mar 25 22:13:30 2021", "2021-04-01"}},
			OtherTypesModel{},
			OtherTypesModel{
				B:          true,
				Time:       time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
				TimeANSIC:  time.Date(2021, time.March, 25, 22, 13, 30, 0, time.UTC),
				TimeFormat: time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		p := NewRowsParser(testCase.RowsCols, testCase.RowsVals, testCase.Model)
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
	testCases := []struct {
		RowsCols []string
		RowsVals [][]string
		Model    interface{}
		Expected interface{}
	}{
		{
			[]string{"int"},
			[][]string{{"-9223372036854775808"}},
			int(0),
			-1 << 63,
		},
		{
			[]string{"int8"},
			[][]string{{"127"}},
			int8(0),
			int8(1<<7 - 1),
		},
		{
			[]string{"int16"},
			[][]string{{"32767"}},
			int16(0),
			int16(1<<15 - 1),
		},
		{
			[]string{"int32"},
			[][]string{{"2147483647"}},
			int32(0),
			int32(1<<31 - 1),
		},
		{
			[]string{"int64"},
			[][]string{{"9223372036854775807"}},
			int64(0),
			int64(1<<63 - 1),
		},
		{
			[]string{"uint"},
			[][]string{{"1"}},
			uint(0),
			uint(1),
		},
		{
			[]string{"uint8"},
			[][]string{{"255"}},
			uint8(0),
			uint8(1<<8 - 1),
		},
		{
			[]string{"uint16"},
			[][]string{{"65535"}},
			uint16(0),
			uint16(1<<16 - 1),
		},
		{
			[]string{"uint32"},
			[][]string{{"4294967295"}},
			uint32(0),
			uint32(1<<32 - 1),
		},
		{
			[]string{"uint64"},
			[][]string{{"18446744073709551615"}},
			uint64(0),
			uint64(1<<64 - 1),
		},
		{
			[]string{"float32"},
			[][]string{{"1.401298464324817070923729583289916131280e-45"}},
			float32(0.0),
			float32(math.SmallestNonzeroFloat32),
		},
		{
			[]string{"float64"},
			[][]string{{"4.940656458412465441765687928682213723651e-324"}},
			float64(0.0),
			float64(math.SmallestNonzeroFloat64),
		},
		{
			[]string{"bool"},
			[][]string{{"true"}},
			bool(false),
			true,
		},
		{
			[]string{"time"},
			[][]string{{"2021-01-02T03:04:05Z"}},
			time.Time{},
			time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, testCase := range testCases {
		p := NewRowsParser(testCase.RowsCols, testCase.RowsVals, testCase.Model)
		p.Next()
		v, err := p.ParseVar(p.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, v.Interface())
	}
}
