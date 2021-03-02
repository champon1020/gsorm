package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	testCases := []struct {
		String string
		Result string
	}{
		{"thisIsString", "this_is_string"},
		{"this", "this"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.Result, internal.SnakeCase(testCase.String))
	}
}

func TestToStringWithQuotes(t *testing.T) {
	var (
		i   int     = 1
		i8  int8    = 2
		i16 int16   = 3
		i32 int32   = 4
		i64 int64   = 5
		u   uint    = 6
		u8  uint8   = 7
		u16 uint16  = 8
		u32 uint32  = 9
		u64 uint64  = 10
		f32 float32 = 10.1
		f64 float64 = 100.1
	)

	testCases := []struct {
		Value  interface{}
		Result string
	}{
		{"rhs", `'rhs'`},
		{i, "1"},
		{i8, "2"},
		{i16, "3"},
		{i32, "4"},
		{i64, "5"},
		{u, "6"},
		{u8, "7"},
		{u16, "8"},
		{u32, "9"},
		{u64, "10"},
		{true, "true"},
		{f32, "10.1"},
		{f64, "100.1"},
		{
			[]interface{}{10, "str", true},
			`10, 'str', true`,
		},
		{
			map[string]string{"key": "value"},
			`map[string]string`,
		},
		{
			nil,
			`nil`,
		},
	}

	for _, testCase := range testCases {
		res := internal.ToString(testCase.Value, true)
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
		res := internal.ToString(testCase.Value, false)
		assert.Equal(t, testCase.Result, res)
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

func TestMapOfColumnsToFields(t *testing.T) {
	type Model1 struct {
		ID        int
		Name      string
		BirthDate time.Time
	}

	type Model2 struct {
		ID   int
		Name string `mgorm:"first_name"`
	}

	testCases := []struct {
		Columns   []string
		ModelType reflect.Type
		Result    map[int]int
	}{
		{
			Columns:   []string{"id", "name", "birth_date"},
			ModelType: reflect.TypeOf(Model1{}),
			Result:    map[int]int{0: 0, 1: 1, 2: 2},
		},
		{
			Columns:   []string{"name", "birth_date", "id"},
			ModelType: reflect.TypeOf(Model1{}),
			Result:    map[int]int{0: 1, 1: 2, 2: 0},
		},
		{
			Columns:   []string{"first_name", "id"},
			ModelType: reflect.TypeOf(Model2{}),
			Result:    map[int]int{0: 1, 1: 0},
		},
		{
			Columns:   []string{"first_name", "first_name", "id"},
			ModelType: reflect.TypeOf(Model2{}),
			Result:    map[int]int{0: 1, 1: 1, 2: 0},
		},
		{
			Columns:   []string{},
			ModelType: reflect.TypeOf(Model2{}),
			Result:    map[int]int{},
		},
	}

	for _, testCase := range testCases {
		res := internal.MapOfColumnsToFields(testCase.Columns, testCase.ModelType)
		assert.Equal(t, testCase.Result, res)
	}
}

func TestColumnNameFromTag(t *testing.T) {
	type Model1 struct {
		UID int `mgorm:"id"`
	}
	m1 := new(Model1)

	type Model2 struct {
		UID int
	}
	m2 := new(Model2)

	type Model3 struct {
		StudentName string
	}
	m3 := new(Model3)

	testCases := []struct {
		Struct reflect.StructField
		Result string
	}{
		{reflect.TypeOf(m1).Elem().Field(0), "id"},
		{reflect.TypeOf(m2).Elem().Field(0), "uid"},
		{reflect.TypeOf(m3).Elem().Field(0), "student_name"},
	}

	for _, testCase := range testCases {
		cn := internal.ColumnNameFromTag(testCase.Struct)
		assert.Equal(t, testCase.Result, cn)
	}
}
