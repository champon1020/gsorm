package mgorm_test

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSQL_String(t *testing.T) {
	testCases := []struct {
		SQL    mgorm.SQL
		Result reflect.Kind
	}{
		{"Test", reflect.String},
	}

	for _, testCase := range testCases {
		sql := mgorm.SQLString(&testCase.SQL)
		assert.Equal(t, testCase.Result, reflect.TypeOf(sql).Kind())
	}
}

func TestSQL_Write(t *testing.T) {
	testCases := []struct {
		SQL    mgorm.SQL
		Str    string
		Result mgorm.SQL
	}{
		{"test", "add", "test add"},
		{"", "add", "add"},
		{"(test", ")", "(test)"},
	}

	for _, testCase := range testCases {
		mgorm.SQLWrite(&testCase.SQL, testCase.Str)
		assert.Equal(t, testCase.Result, testCase.SQL)
	}
}

func TestSQL_DoQuery(t *testing.T) {
	type Car struct {
		ID   int `mgorm:id`
		Name string
	}

	testCases := []struct {
		Rows *[]Car
	}{
		{&[]Car{{ID: 100, Name: "test"}}},
		{&[]Car{{ID: 100, Name: "test"}, {ID: 200, Name: "test2"}}},
	}

	s := new(mgorm.SQL)
	for _, testCase := range testCases {
		car := new([]Car)
		mockRows := new(mgorm.TestMockRows)
		mockRows.Max = len(*testCase.Rows)
		mockRows.ColumnsFunc = func() ([]string, error) { return []string{"id", "name"}, nil }
		mockRows.ScanFunc = func(dest ...interface{}) error {
			ptrID := dest[0].(*[]byte)
			ptrName := dest[1].(*[]byte)
			*ptrID = []byte(fmt.Sprintf("%d", (*testCase.Rows)[mockRows.Count-1].ID))
			*ptrName = []byte((*testCase.Rows)[mockRows.Count-1].Name)
			return nil
		}
		mockdb := &mgorm.TestMockDB{
			QueryFunc: func(string, ...interface{}) (mgorm.Rows, error) { return mockRows, nil },
		}
		if err := mgorm.SQLDoQuery(s, mockdb, car); err != nil {
			t.Error(err)
		}
		if diff := cmp.Diff(car, testCase.Rows); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestSQL_DoQuery_Fail(t *testing.T) {
	type Model struct{}

	testCases := []struct {
		Model     interface{}
		QueryFunc func(string, ...interface{}) (mgorm.Rows, error)
		Error     error
	}{
		{
			&[]Model{},
			func(string, ...interface{}) (mgorm.Rows, error) { return nil, errors.New("test1") },
			internal.NewError(mgorm.OpSQLDoQuery, internal.KindDatabase, errors.New("test1")),
		},
		{
			&[]Model{},
			func(string, ...interface{}) (mgorm.Rows, error) {
				return &mgorm.TestMockRows{
					Max:         1,
					ColumnsFunc: func() ([]string, error) { return []string{}, errors.New("test2") },
					ScanFunc:    func(...interface{}) error { return nil },
				}, nil
			},
			internal.NewError(mgorm.OpSQLDoQuery, internal.KindDatabase, errors.New("test2")),
		},
		{
			&Model{},
			func(string, ...interface{}) (mgorm.Rows, error) {
				return &mgorm.TestMockRows{
					Max:         1,
					ColumnsFunc: func() ([]string, error) { return []string{}, nil },
					ScanFunc:    func(...interface{}) error { return nil },
				}, nil
			},
			internal.NewError(
				mgorm.OpSQLDoQuery,
				internal.KindType,
				errors.New("model type must be slice or array"),
			),
		},
		{
			&[]Model{},
			func(string, ...interface{}) (mgorm.Rows, error) {
				return &mgorm.TestMockRows{
					Max:         1,
					ColumnsFunc: func() ([]string, error) { return []string{}, nil },
					ScanFunc:    func(...interface{}) error { return errors.New("test3") },
				}, nil
			},
			internal.NewError(mgorm.OpSQLDoQuery, internal.KindDatabase, errors.New("test3")),
		},
	}

	s := new(mgorm.SQL)
	for _, testCase := range testCases {
		mockdb := &mgorm.TestMockDB{QueryFunc: testCase.QueryFunc}
		err := mgorm.SQLDoQuery(s, mockdb, testCase.Model)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(*e, *testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestSQL_DoExec(t *testing.T) {
	s := new(mgorm.SQL)
	flg := false
	mockdb := &mgorm.TestMockDB{ExecFunc: func(string, ...interface{}) (sql.Result, error) {
		flg = true
		return nil, nil
	}}
	if err := mgorm.SQLDoExec(s, mockdb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, flg)
}

func TestSQL_DoExec_Fail(t *testing.T) {
	type Model struct{}

	testCases := []struct {
		ExecFunc func(string, ...interface{}) (sql.Result, error)
		Error    error
	}{
		{
			func(string, ...interface{}) (sql.Result, error) { return nil, errors.New("test") },
			internal.NewError(mgorm.OpSQLDoExec, internal.KindDatabase, errors.New("test")),
		},
	}

	s := new(mgorm.SQL)
	for _, testCase := range testCases {
		mockdb := &mgorm.TestMockDB{ExecFunc: testCase.ExecFunc}
		err := mgorm.SQLDoExec(s, mockdb)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(*e, *testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
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
		cn := mgorm.ColumnName(testCase.Struct)
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
		Value    []byte
		Result   reflect.Value
	}{
		{
			0,
			[]byte("test"),
			reflect.ValueOf(&Car{Name: "test"}),
		},
		{
			1,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{ID: 100}),
		},
		{
			2,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{ID8: 100}),
		},
		{
			3,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{ID16: 100}),
		},
		{
			4,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{ID32: 100}),
		},
		{
			5,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{ID64: 100}),
		},
		{
			6,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{UID: 100}),
		},
		{
			7,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{UID8: 100}),
		},
		{
			8,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{UID16: 100}),
		},
		{
			9,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{UID32: 100}),
		},
		{
			10,
			[]byte(fmt.Sprintf("%d", 100)),
			reflect.ValueOf(&Car{UID64: 100}),
		},
		{
			11,
			[]byte(fmt.Sprintf("%f", 100.2)),
			reflect.ValueOf(&Car{FID32: 100.2}),
		},
		{
			12,
			[]byte(fmt.Sprintf("%f", 100.2)),
			reflect.ValueOf(&Car{FID64: 100.2}),
		},
		{
			13,
			[]byte("true"),
			reflect.ValueOf(&Car{Flg: true}),
		},
		{
			14,
			[]byte("2020-01-02T03:04:05Z"),
			reflect.ValueOf(&Car{
				Time: time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
			}),
		},
		{
			15,
			[]byte("2020-01-02"),
			reflect.ValueOf(&Car{
				Time2: time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			}),
		},
		{
			16,
			[]byte("Thu Jan 2 03:04:05 2020"),
			reflect.ValueOf(&Car{
				Time3: time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC),
			}),
		},
	}

	for _, testCase := range testCases {
		v := reflect.ValueOf(new(Car))
		sf := reflect.TypeOf(Car{}).Field(testCase.FieldNum)
		if err := mgorm.SetField(reflect.Indirect(v).Field(testCase.FieldNum), sf, testCase.Value); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.Result.Interface(), v.Interface())
	}
}
