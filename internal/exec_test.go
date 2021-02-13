package internal_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

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

func TestSetValueToField(t *testing.T) {
	type Car struct {
		Name  string
		ID    int
		ID8   int8
		ID16  int16
		ID32  int32
		ID64  int64
		UID   uint
		UID8  uint8
		UID16 uint16
		UID32 uint32
		UID64 uint64
		FID32 float32
		FID64 float64
		Flg   bool
		Time  time.Time
		Time2 time.Time `layout:"2006-01-02"`
		Time3 time.Time `layout:"time.ANSIC"`
	}

	testCases := []struct {
		FieldIndex int
		Value      string
		Result     reflect.Value
	}{
		{
			0,
			"test",
			reflect.ValueOf(&Car{Name: "test"}),
		},
		{
			1,
			fmt.Sprintf("%d", 10),
			reflect.ValueOf(&Car{ID: 10}),
		},
		{
			2,
			fmt.Sprintf("%d", 10),
			reflect.ValueOf(&Car{ID8: 10}),
		},
		{
			3,
			fmt.Sprintf("%d", 10),
			reflect.ValueOf(&Car{ID16: 10}),
		},
		{
			4,
			fmt.Sprintf("%d", 10),
			reflect.ValueOf(&Car{ID32: 10}),
		},
		{
			5,
			fmt.Sprintf("%d", 10),
			reflect.ValueOf(&Car{ID64: 10}),
		},
		{
			6,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{UID: 100}),
		},
		{
			7,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{UID8: 100}),
		},
		{
			8,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{UID16: 100}),
		},
		{
			9,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{UID32: 100}),
		},
		{
			10,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{UID64: 100}),
		},
		{
			11,
			fmt.Sprintf("%f", 100.2),
			reflect.ValueOf(&Car{FID32: 100.2}),
		},
		{
			12,
			fmt.Sprintf("%f", 100.2),
			reflect.ValueOf(&Car{FID64: 100.2}),
		},
		{
			13,
			"true",
			reflect.ValueOf(&Car{Flg: true}),
		},
		{
			14,
			"2020-01-02T03:04:05Z",
			reflect.ValueOf(&Car{
				Time: time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
			}),
		},
		{
			15,
			"2020-01-02",
			reflect.ValueOf(&Car{
				Time2: time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			}),
		},
		{
			16,
			"Thu Jan 2 03:04:05 2020",
			reflect.ValueOf(&Car{
				Time3: time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
			}),
		},
	}

	for _, testCase := range testCases {
		ref := reflect.ValueOf(new(Car))
		if err := internal.SetValueToField(
			ref.Elem(),
			testCase.FieldIndex,
			testCase.Value,
		); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Result.Interface(), ref.Interface())
	}
}

func TestSetValueToMap(t *testing.T) {
	m := make(map[int]string)

	testCases := []struct {
		Ref    interface{}
		Key    string
		Value  string
		Result reflect.Value
	}{
		{
			Ref:    &m,
			Key:    "10",
			Value:  "str",
			Result: reflect.ValueOf(&map[int]string{10: "str"}),
		},
	}

	for _, testCase := range testCases {
		mResult := make(map[int]string)
		ref := reflect.ValueOf(&mResult)
		if err := internal.SetValueToMap(ref.Elem(), testCase.Key, testCase.Value); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result.Interface(), ref.Interface()); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestSetValueToVar(t *testing.T) {
	var (
		s   string
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
		u   uint
		u8  uint8
		u16 uint16
		u32 uint32
		u64 uint64
		f32 float32
		f64 float64
		b   bool
		tm  time.Time
	)

	testCases := []struct {
		Ref    interface{}
		Value  string
		Result reflect.Value
	}{
		{
			Ref:    &s,
			Value:  "str",
			Result: reflect.ValueOf("str"),
		},
		{
			Ref:    &i,
			Value:  "10",
			Result: reflect.ValueOf(10),
		},
		{
			Ref:    &i8,
			Value:  "10",
			Result: reflect.ValueOf(int8(10)),
		},
		{
			Ref:    &i16,
			Value:  "10",
			Result: reflect.ValueOf(int16(10)),
		},
		{
			Ref:    &i32,
			Value:  "10",
			Result: reflect.ValueOf(int32(10)),
		},
		{
			Ref:    &i64,
			Value:  "10",
			Result: reflect.ValueOf(int64(10)),
		},
		{
			Ref:    &u,
			Value:  "100",
			Result: reflect.ValueOf(uint(100)),
		},
		{
			Ref:    &u8,
			Value:  "100",
			Result: reflect.ValueOf(uint8(100)),
		},
		{
			Ref:    &u16,
			Value:  "100",
			Result: reflect.ValueOf(uint16(100)),
		},
		{
			Ref:    &u32,
			Value:  "100",
			Result: reflect.ValueOf(uint32(100)),
		},
		{
			Ref:    &u64,
			Value:  "100",
			Result: reflect.ValueOf(uint64(100)),
		},
		{
			Ref:    &f32,
			Value:  "100.2",
			Result: reflect.ValueOf(float32(100.2)),
		},
		{
			Ref:    &f64,
			Value:  "100.2",
			Result: reflect.ValueOf(float64(100.2)),
		},
		{
			Ref:    &b,
			Value:  "true",
			Result: reflect.ValueOf(true),
		},
		{
			Ref:    &tm,
			Value:  "2020-01-02T03:04:05Z",
			Result: reflect.ValueOf(time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC)),
		},
	}

	for _, testCase := range testCases {
		ref := reflect.ValueOf(testCase.Ref).Elem()
		if err := internal.SetValueToVar(ref, testCase.Value); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		assert.Equal(t, testCase.Result.Interface(), ref.Interface())
	}
}
