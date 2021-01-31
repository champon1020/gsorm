package internal_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

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

func TestSetField(t *testing.T) {
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
		FieldNum int
		Value    string
		Result   reflect.Value
	}{
		{
			0,
			"test",
			reflect.ValueOf(&Car{Name: "test"}),
		},
		{
			1,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{ID: 100}),
		},
		{
			2,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{ID8: 100}),
		},
		{
			3,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{ID16: 100}),
		},
		{
			4,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{ID32: 100}),
		},
		{
			5,
			fmt.Sprintf("%d", 100),
			reflect.ValueOf(&Car{ID64: 100}),
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
		v := reflect.ValueOf(new(Car))
		sf := reflect.TypeOf(Car{}).Field(testCase.FieldNum)
		if err := internal.SetField(reflect.Indirect(v).Field(testCase.FieldNum), sf, testCase.Value); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.Result.Interface(), v.Interface())
	}
}
