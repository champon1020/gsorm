package internal_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

func TestConvertToSnakeCase(t *testing.T) {
	testCases := []struct {
		String string
		Result string
	}{
		{"thisIsString", "this_is_string"},
		{"this", "this"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.Result, internal.ConvertToSnakeCase(testCase.String))
	}
}

func TestToStringWithQuotes(t *testing.T) {
	var (
		n0 int    = 1
		n1 int8   = 2
		n2 int16  = 3
		n3 int32  = 4
		n4 int64  = 5
		u0 uint   = 6
		u1 uint8  = 7
		u2 uint16 = 8
		u3 uint32 = 9
		u4 uint64 = 10
	)

	testCases := []struct {
		Value  interface{}
		Result string
	}{
		{"rhs", `"rhs"`},
		{n0, "1"},
		{n1, "2"},
		{n2, "3"},
		{n3, "4"},
		{n4, "5"},
		{u0, "6"},
		{u1, "7"},
		{u2, "8"},
		{u3, "9"},
		{u4, "10"},
		{true, "true"},
	}

	for _, testCase := range testCases {
		res, _ := internal.ToString(testCase.Value, true)
		assert.Equal(t, testCase.Result, res)
	}
}

func TestToStringWituhoutQuotes(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result string
	}{
		{"rhs", "rhs"},
	}

	for _, testCase := range testCases {
		res, _ := internal.ToString(testCase.Value, false)
		assert.Equal(t, testCase.Result, res)
	}
}

func TestToString_Fail(t *testing.T) {
	testCases := []struct {
		Value interface{}
		Error error
	}{
		{
			map[string]string{"key": "value"},
			internal.NewError(internal.OpToString, internal.KindType, errors.New("type is invalid")),
		},
		{
			[]int{1, 2},
			internal.NewError(internal.OpToString, internal.KindType, errors.New("type is invalid")),
		},
		{
			[2]int{1, 2},
			internal.NewError(internal.OpToString, internal.KindType, errors.New("type is invalid")),
		},
	}

	for _, testCase := range testCases {
		_, err := internal.ToString(testCase.Value, false)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}
		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		assert.Equal(t, *e, *testCase.Error.(*internal.Error))
	}
}

func TestSliceToString(t *testing.T) {
	testCases := []struct {
		Values []interface{}
		Result string
	}{
		{
			[]interface{}{10, "str", true},
			`10, "str", true`,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.Result, internal.SliceToString(testCase.Values))
	}
}

func TestMapKeyType(t *testing.T) {
	var (
		m0 = make(map[int]uint)
		m1 = make(map[int8]int)
		m2 = make(map[int16]int8)
		m3 = make(map[int32]int16)
		m4 = make(map[int64]int32)
		m5 = make(map[uint]int64)
	)

	testCases := []struct {
		MapRef interface{}
		Result reflect.Type
	}{
		{m0, reflect.TypeOf(int(0))},
		{m1, reflect.TypeOf(int8(0))},
		{m2, reflect.TypeOf(int16(0))},
		{m3, reflect.TypeOf(int32(0))},
		{m4, reflect.TypeOf(int64(0))},
		{m5, reflect.TypeOf(uint(0))},
	}

	for _, testCase := range testCases {
		typ := internal.MapKeyType(reflect.TypeOf(testCase.MapRef))
		assert.Equal(t, testCase.Result, typ)
	}
}

func TestMapValueType(t *testing.T) {
	var (
		m0 = make(map[uint64]uint8)
		m1 = make(map[uint8]uint16)
		m2 = make(map[uint16]uint32)
		m3 = make(map[uint32]uint64)
		m4 = make(map[float64]float32)
		m5 = make(map[float32]float64)
		m6 = make(map[string]bool)
		m7 = make(map[bool]string)
	)

	testCases := []struct {
		MapRef interface{}
		Result reflect.Type
	}{
		{m0, reflect.TypeOf(uint8(0))},
		{m1, reflect.TypeOf(uint16(0))},
		{m2, reflect.TypeOf(uint32(0))},
		{m3, reflect.TypeOf(uint64(0))},
		{m4, reflect.TypeOf(float32(0.0))},
		{m5, reflect.TypeOf(float64(0.0))},
		{m6, reflect.TypeOf(false)},
		{m7, reflect.TypeOf("")},
	}

	for _, testCase := range testCases {
		typ := internal.MapValueType(reflect.TypeOf(testCase.MapRef))
		assert.Equal(t, testCase.Result, typ)
	}
}
