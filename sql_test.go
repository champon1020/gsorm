package mgorm_test

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSQL_String(t *testing.T) {
	testCases := []struct {
		SQL    *SQL
		Result reflect.Kind
	}{
		{"Test", reflect.String},
	}

	for _, testCase := range testCases {
		sql := SQLString(testCase.SQL)
		assert.Equal(t, testCase.Result, reflect.TypeOf(sql).Kind())
	}
}

func TestSQL_Write(t *testing.T) {
	testCases := []struct {
		SQL    *SQL
		Str    string
		Result SQL
	}{
		{"test", "add", "test add"},
		{"", "add", "add"},
		{"(test", ")", "(test)"},
	}

	for _, testCase := range testCases {
		SQLWrite(testCase.SQL, testCase.Str)
		assert.Equal(t, testCase.Result, testCase.SQL)
	}
}

func TestSQL_DoExec(t *testing.T) {
	s := new(SQL)
	flg := false
	mockdb := &MockDb{ExecFunc: func(string, ...interface{}) (sql.Result, error) {
		flg = true
		return nil, nil
	}}
	if err := SQLDoExec(s, mockdb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, flg)
}

func TestSQL_DoExec_Fail(t *testing.T) {
	type Model struct{}

	testCases := []struct {
		ExecFunc  func(string, ...interface{}) (sql.Result, error)
		ErrorCode int
	}{
		{
			func(string, ...interface{}) (sql.Result, error) { return nil, errors.New("") },
			ErrExecFailed,
		},
	}

	s := new(SQL)
	for _, testCase := range testCases {
		mockdb := &MockDb{ExecFunc: testCase.ExecFunc}
		err := SQLDoExec(s, mockdb)
		if err == nil {
			t.Errorf("error is nil, %v", testCase)
		}
		assert.Equal(t, testCase.ErrorCode, err.(Error).Code)
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
		cn := ColumnName(testCase.Struct)
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
		if err := SetField(reflect.Indirect(v).Field(testCase.FieldNum), testCase.Value); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.Result.Interface(), v.Interface())
	}
}
