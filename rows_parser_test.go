package gsorm_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

type fakeDB struct {
	r gsorm.ExportedIRows
}

func newFakeDB(r gsorm.ExportedIRows) gsorm.DB {
	return &fakeDB{r: r}
}

func (d *fakeDB) Ping() error {
	return nil
}

func (d *fakeDB) Query(query string, args ...interface{}) (gsorm.ExportedIRows, error) {
	return d.r, nil
}

func (d *fakeDB) Exec(query string, args ...interface{}) (gsorm.ExportedIResult, error) {
	return nil, nil
}

func (d *fakeDB) SetConnMaxLifetime(n time.Duration) error {
	return nil
}

func (d *fakeDB) SetMaxIdleConns(n int) error {
	return nil
}

func (d *fakeDB) SetMaxOpenConns(n int) error {
	return nil
}

func (d *fakeDB) Close() error {
	return nil
}

func (d *fakeDB) Begin() (gsorm.Tx, error) {
	return nil, nil
}

type fakeRows struct {
	ct  []gsorm.ExportedIColumnType
	v   [][]interface{}
	itr int
}

func newFakeRows(ct []gsorm.ExportedIColumnType, v [][]interface{}) gsorm.ExportedIRows {
	return &fakeRows{ct: ct, v: v, itr: -1}
}

func (r *fakeRows) Next() bool {
	if r.itr+1 < len(r.v) {
		r.itr++
		return true
	}
	return false
}

func (r *fakeRows) Scan(args ...interface{}) error {
	for i, a := range args {
		v := reflect.ValueOf(r.v[r.itr][i])
		reflect.ValueOf(a).Elem().Set(v)
	}
	return nil
}

func (r *fakeRows) ColumnTypes() ([]gsorm.ExportedIColumnType, error) {
	return r.ct, nil
}

func (r *fakeRows) Close() error {
	return nil
}

type fakeColumnType struct {
	n string
	t reflect.Type
}

func newFakeColumn(n string, t reflect.Type) gsorm.ExportedIColumnType {
	return &fakeColumnType{n: n, t: t}
}

func (c *fakeColumnType) Name() string {
	return c.n
}

func (c *fakeColumnType) ScanType() reflect.Type {
	return c.t
}

func TestRowsParser_ParseMapSlice(t *testing.T) {
	model := []map[string]interface{}{}
	expected := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
		{"emp_no": 1003, "first_name": "Saburo"},
	}

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("emp_no", reflect.TypeOf(0)),
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	var v [][]interface{} = make([][]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		v[i] = append(v[i], expected[i]["emp_no"])
		v[i] = append(v[i], expected[i]["first_name"])
	}
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "emp_no", "first_name").From("employees").Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	if diff := cmp.Diff(expected, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRowsParser_ParseStructSlice(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}

	model := []Employee{}
	expected := []Employee{
		{ID: 1001, FirstName: "Taro"},
		{ID: 1002, FirstName: "Jiro"},
		{ID: 1003, FirstName: "Saburo"},
	}

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("emp_no", reflect.TypeOf(0)),
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	v := make([][]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		v[i] = append(v[i], expected[i].ID)
		v[i] = append(v[i], expected[i].FirstName)
	}
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "emp_no", "first_name").From("employees").Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	if diff := cmp.Diff(expected, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRowsParser_ParseSlice(t *testing.T) {
	model := []string{}
	expected := []string{"Taro", "Jiro", "Saburo"}

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	v := make([][]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		v[i] = append(v[i], expected[i])
	}
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "first_name").From("employees").Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	if diff := cmp.Diff(expected, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRowsParser_ParseMap(t *testing.T) {
	model := map[string]interface{}{}
	expected := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("emp_no", reflect.TypeOf(0)),
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	v := make([][]interface{}, 1)
	v[0] = append(v[0], expected["emp_no"])
	v[0] = append(v[0], expected["first_name"])
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "emp_no", "first_name").From("employees").Limit(1).Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	if diff := cmp.Diff(expected, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRowsParser_ParseStruct(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}

	model := Employee{}
	expected := Employee{ID: 1001, FirstName: "Taro"}

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("emp_no", reflect.TypeOf(0)),
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	v := make([][]interface{}, 1)
	v[0] = append(v[0], expected.ID)
	v[0] = append(v[0], expected.FirstName)
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "emp_no", "first_name").From("employees").Limit(1).Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	if diff := cmp.Diff(expected, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRowsParser_ParseVar(t *testing.T) {
	var model string
	expected := "Taro"

	// Prepare the fake connection.
	ct := []gsorm.ExportedIColumnType{
		newFakeColumn("first_name", reflect.TypeOf("")),
	}
	v := make([][]interface{}, 1)
	v[0] = append(v[0], expected)
	rows := newFakeRows(ct, v)
	db := newFakeDB(rows)

	// Actual process.
	if err := gsorm.Select(db, "first_name").From("employees").Limit(1).Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	// Validate.
	assert.Equal(t, expected, model)
}
