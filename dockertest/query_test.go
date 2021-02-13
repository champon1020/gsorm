package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestQueryWithSlice(t *testing.T) {
	// SELECT emp_no FROM employees;
	stmt := mgorm.Select(db, "emp_no").From("employees")
	result := &[]int{10001, 10002, 10003, 10004, 10005, 10006, 10007, 10008, 10009, 10010}

	model := new([]int)
	if err := stmt.Query(model); err != nil {
		t.Errorf("Error was occurred: %v", err)
		t.Errorf("Executed SQL: %s", stmt.String())
	}
	if diff := cmp.Diff(result, model); diff != "" {
		t.Errorf("Executed SQL: %s", stmt.String())
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestQueryWithVar(t *testing.T) {
	// SELECT COUNT(emp_no) FROM employees;
	stmt := mgorm.Count(db, "emp_no").From("employees")
	result := 10

	model := new(int)
	if err := stmt.Query(model); err != nil {
		t.Errorf("Error was occurred: %v", err)
		t.Errorf("Executed SQL: %s", stmt.String())
	}
	assert.Equal(t, result, *model)
}

func TestQueryWithStruct(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *Employee
	}{
		// SELECT emp_no FROM employees LIMIT 1;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Limit(1).(*mgorm.Stmt),
			&Employee{EmpNo: 10001},
		},

		// SELECT emp_no FROM employees;
		{
			mgorm.Select(db, "emp_no").From("employees").(*mgorm.Stmt),
			&Employee{EmpNo: 10001},
		},
	}

	for i, testCase := range testCases {
		model := new(Employee)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Executed SQL: %s", testCase.Stmt.String())
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestQueryWithMap(t *testing.T) {
	stmt := mgorm.Select(db, "emp_no", "first_name").From("employees")
	result := map[int]string{
		10001: "Georgi",
		10002: "Bezalel",
		10003: "Parto",
		10004: "Chirstian",
		10005: "Kyoichi",
		10006: "Anneke",
		10007: "Tzvetan",
		10008: "Saniya",
		10009: "Sumant",
		10010: "Duangkaew",
	}

	model := make(map[int]string)
	if err := stmt.Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
		t.Errorf("Executed SQL: %s", stmt.String())
	}
	if diff := cmp.Diff(result, model); diff != "" {
		t.Errorf("Executed SQL: %s", stmt.String())
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}
