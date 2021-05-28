package parser_test

import (
	"math"
	"testing"
	"time"

	"github.com/champon1020/gsorm/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestUpdateModelParser_ParseMap(t *testing.T) {
	testCases := []struct {
		Cols     []string
		Model    interface{}
		Expected string
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
			&map[string]interface{}{
				"int":   -9223372036854775808,
				"int8":  127,
				"int16": 32767,
				"int32": 2147483647,
				"int64": 9223372036854775807,
			},
			"int = -9223372036854775808, int8 = 127, int16 = 32767, int32 = 2147483647, int64 = 9223372036854775807",
		},
		{
			[]string{"float32", "float64"},
			&map[string]interface{}{
				"float32": float32(math.Pi),
				"float64": float64(math.Pi),
			},
			"float32 = 3.1415927, float64 = 3.141592653589793",
		},
		{
			[]string{"bool", "time"},
			&map[string]interface{}{
				"bool": true,
				"time": time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
			},
			"bool = 1, time = '2021-01-02 03:04:05'",
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewUpdateModelParser(testCase.Cols, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		sql, err := p.Parse()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, sql.String())
	}
}

func TestUpdateModelParser_ParseStruct(t *testing.T) {
	testCases := []struct {
		Cols     []string
		Model    interface{}
		Expected string
	}{
		{
			[]string{"int", "int8", "int16", "int32", "int64"},
			&IntModel{I8: 127, I16: 32767, I32: 2147483647, I64: 9223372036854775807, I: -9223372036854775808},
			"int = -9223372036854775808, int8 = 127, int16 = 32767, int32 = 2147483647, int64 = 9223372036854775807",
		},
		{
			[]string{"uint", "uint8", "uint16", "uint32", "uint64"},
			&UintModel{U8: 255, U16: 65535, U32: 4294967295, U64: 18446744073709551615, U: 1},
			"uint = 1, uint8 = 255, uint16 = 65535, uint32 = 4294967295, uint64 = 18446744073709551615",
		},
		{
			[]string{"float32", "float64"},
			&FloatModel{F32: math.Pi, F64: math.Pi},
			"float32 = 3.1415927, float64 = 3.141592653589793",
		},
		{
			[]string{"bool", "time"},
			&OtherTypesModel{
				B:    true,
				Time: time.Date(2021, time.January, 2, 3, 4, 5, 0, time.UTC),
			},
			"bool = 1, time = '2021-01-02 03:04:05'",
		},
	}

	for _, testCase := range testCases {
		p, err := parser.NewUpdateModelParser(testCase.Cols, testCase.Model)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		sql, err := p.Parse()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Expected, sql.String())
	}
}
