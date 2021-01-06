package syntax

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSQL_String(t *testing.T) {
}

func TestSQL_Write(t *testing.T) {
}

func TestSQL_DoQuery(t *testing.T) {
}

func TestSQL_DoExec(t *testing.T) {
}

func TestSetRowsToModel(t *testing.T) {
	type Car struct {
		Id   int
		Name string
	}

	c := new([]Car)
	testCases := []struct {
		Mv     reflect.Value
		Mt     reflect.Type
		cols   []string
		rowVal []interface{}
	}{
		{
			reflect.New(reflect.TypeOf(c)),
			reflect.TypeOf(c),
			[]string{"Id", "Name"},
			[]interface{}{1, "test"},
		},
	}

	for testCase := range testCases {

	}
}

func TestColumnName(t *testing.T) {
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
	}

	testCases := []struct {
		FieldNum int
		Value    interface{}
		Result   reflect.Value
	}{
		{
			0,
			"test",
			reflect.ValueOf(&Car{Name: "test"}),
		},
		{
			1,
			100,
			reflect.ValueOf(&Car{ID: 100}),
		},
		{
			2,
			100,
			reflect.ValueOf(&Car{ID8: 100}),
		},
		{
			3,
			int16(100),
			reflect.ValueOf(&Car{ID16: 100}),
		},
		{
			4,
			int64(100),
			reflect.ValueOf(&Car{ID32: 100}),
		},
		{
			5,
			int32(100),
			reflect.ValueOf(&Car{ID64: 100}),
		},
		{
			6,
			100,
			reflect.ValueOf(&Car{UID: 100}),
		},
		{
			7,
			100,
			reflect.ValueOf(&Car{UID8: 100}),
		},
		{
			8,
			int16(100),
			reflect.ValueOf(&Car{UID16: 100}),
		},
		{
			9,
			int64(100),
			reflect.ValueOf(&Car{UID32: 100}),
		},
		{
			10,
			int32(100),
			reflect.ValueOf(&Car{UID64: 100}),
		},
		{
			11,
			float64(100.2),
			reflect.ValueOf(&Car{FID32: 100.2}),
		},
		{
			12,
			float32(100.2),
			reflect.ValueOf(&Car{FID64: 100.2}),
		},
		{
			13,
			true,
			reflect.ValueOf(&Car{Flg: true}),
		},
		{
			14,
			time.Date(2020, time.January, 2, 3, 4, 5, 6, time.UTC),
			reflect.ValueOf(&Car{Time: time.Date(2020, time.January, 2, 3, 4, 5, 6, time.UTC)}),
		},
	}

	for _, testCase := range testCases {
		v := reflect.ValueOf(new(Car))
		if err := setField(reflect.Indirect(v).Field(testCase.FieldNum), testCase.Value); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.Result.Interface(), v.Interface())
	}
}
