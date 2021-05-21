package e2e_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestQueryWithSlice(t *testing.T) {
	// SELECT emp_no FROM employees;
	stmt := gsorm.Select(db, "emp_no").From("employees")
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
	stmt := gsorm.Count(db, "emp_no").From("employees")
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
		Stmt   *statement.SelectStmt
		Result *Employee
	}{
		// SELECT emp_no FROM employees LIMIT 1;
		{
			gsorm.Select(db, "emp_no").
				From("employees").
				Limit(1).(*statement.SelectStmt),
			&Employee{EmpNo: 10001},
		},

		// SELECT emp_no FROM employees;
		{
			gsorm.Select(db, "emp_no").From("employees").(*statement.SelectStmt),
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
	stmt := gsorm.Select(db, "emp_no", "first_name").From("employees")
	result := []map[string]interface{}{
		{
			"emp_no":     10001,
			"first_name": "Georgi",
		},
		{
			"emp_no":     10002,
			"first_name": "Bezalel",
		},
		{
			"emp_no":     10003,
			"first_name": "Parto",
		},
		{
			"emp_no":     10004,
			"first_name": "Chirstian",
		},
		{
			"emp_no":     10005,
			"first_name": "Kyoichi",
		},
		{
			"emp_no":     10006,
			"first_name": "Anneke",
		},
		{
			"emp_no":     10007,
			"first_name": "Tzvetan",
		},
		{
			"emp_no":     10008,
			"first_name": "Saniya",
		},
		{
			"emp_no":     10009,
			"first_name": "Sumant",
		},
		{
			"emp_no":     10010,
			"first_name": "Duangkaew",
		},
	}

	model := make([]map[string]interface{}, 10)
	if err := stmt.Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
		t.Errorf("Executed SQL: %s", stmt.String())
	}
	if diff := cmp.Diff(result, model); diff != "" {
		t.Errorf("Executed SQL: %s", stmt.String())
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}
