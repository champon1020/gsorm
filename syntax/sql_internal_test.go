package syntax

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSQL_String(t *testing.T) {
	testCases := []struct {
		SQL    SQL
		Result reflect.Kind
	}{
		{"Test", reflect.String},
	}

	for _, testCase := range testCases {
		sql := testCase.SQL.string()
		assert.Equal(t, testCase.Result, reflect.TypeOf(sql).Kind())
	}
}

func TestSQL_Write(t *testing.T) {
	testCases := []struct {
		SQL    SQL
		Str    string
		Result SQL
	}{
		{"test", "add", "test add"},
		{"", "add", "add"},
		{"(test", ")", "(test)"},
	}

	for _, testCase := range testCases {
		testCase.SQL.write(testCase.Str)
		assert.Equal(t, testCase.Result, testCase.SQL)
	}
}

type MockDb struct {
	QueryFunc func(string, ...interface{}) (RowsIface, error)
	ExecFunc  func(string, ...interface{}) (sql.Result, error)
}

func (db *MockDb) Query(query string, args ...interface{}) (RowsIface, error) {
	return db.QueryFunc(query, args...)
}
func (db *MockDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.ExecFunc(query, args...)
}

type MockRows struct {
	Max      int
	Count    int
	ScanFunc func(...interface{}) error
}

func (r *MockRows) Close() error { return nil }
func (r *MockRows) Columns() ([]string, error) {
	return []string{"id", "name"}, nil
}
func (r *MockRows) Next() bool {
	if r.Count >= r.Max {
		return false
	}
	r.Count++
	return true
}
func (r *MockRows) Scan(dest ...interface{}) error {
	return r.ScanFunc(dest...)
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

	s := new(SQL)
	for _, testCase := range testCases {
		car := new([]Car)
		mockRows := new(MockRows)
		mockRows.Max = len(*testCase.Rows)
		mockRows.ScanFunc = func(dest ...interface{}) error {
			ptrID := dest[0].(*interface{})
			ptrName := dest[1].(*interface{})
			*ptrID = (*testCase.Rows)[mockRows.Count-1].ID
			*ptrName = (*testCase.Rows)[mockRows.Count-1].Name
			return nil
		}
		mockdb := &MockDb{QueryFunc: func(string, ...interface{}) (RowsIface, error) {
			return mockRows, nil
		}}
		if err := s.doQuery(mockdb, car); err != nil {
			t.Error(err)
		}
		if diff := cmp.Diff(car, testCase.Rows); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}

func TestSQL_DoExec(t *testing.T) {
	s := new(SQL)
	flg := false
	mockdb := &MockDb{ExecFunc: func(string, ...interface{}) (sql.Result, error) {
		flg = true
		return nil, nil
	}}
	if err := s.doExec(mockdb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, true, flg)
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
		cn := columnName(testCase.Struct)
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
		if err := setField(reflect.Indirect(v).Field(testCase.FieldNum), testCase.Value); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.Result.Interface(), v.Interface())
	}
}
