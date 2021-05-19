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
		{"ID", "id"},
		{"UUID", "uuid"},
		{"EmpID", "emp_id"},
		{"HTTPRequest", "httprequest"},
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
		{true, "1"},
		{false, "0"},
		{f32, "10.1"},
		{f64, "100.1"},
		{
			[]interface{}{10, "str", true},
			`10, 'str', 1`,
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
		res := internal.ToString(testCase.Value, nil)
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
		res := internal.ToString(testCase.Value, &internal.ToStringOpt{Quotes: false})
		assert.Equal(t, testCase.Result, res)
	}
}

func TestColumnsAndFields(t *testing.T) {
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
		res := internal.ColumnsAndFields(testCase.Columns, testCase.ModelType)
		assert.Equal(t, testCase.Result, res)
	}
}

func TestColumnName(t *testing.T) {
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
		cn := internal.ColumnName(testCase.Struct)
		assert.Equal(t, testCase.Result, cn)
	}
}
