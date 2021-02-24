package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/google/go-cmp/cmp"
)

func TestMaxMin(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.SelectStmt
		Result *Employee
	}{
		// SELECT MIN(emp_no) FROM employees;
		{
			mgorm.Min(db, "emp_no").
				From("employees").(*mgorm.SelectStmt),
			&Employee{
				EmpNo: 10001,
			},
		},

		// SELECT MAX(emp_no) FROM employees;
		{
			mgorm.Max(db, "emp_no").
				From("employees").(*mgorm.SelectStmt),
			&Employee{
				EmpNo: 10010,
			},
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

func TestCountSum(t *testing.T) {
	var (
		cnt = 10
		sum = 579290
	)

	testCases := []struct {
		Stmt   *mgorm.SelectStmt
		Result *int
	}{
		// SELECT COUNT(emp_no) FROM salaries;
		{
			mgorm.Count(db, "emp_no").
				From("salaries").(*mgorm.SelectStmt),
			&cnt,
		},

		// SELECT SUM(salary) FROM salaries;
		{
			mgorm.Sum(db, "salary").
				From("salaries").(*mgorm.SelectStmt),
			&sum,
		},
	}

	for i, testCase := range testCases {
		model := new(int)
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

func TestAvg(t *testing.T) {
	avg := 57929.0

	testCases := []struct {
		Stmt   *mgorm.SelectStmt
		Result *float64
	}{
		// SELECT AVG(salary) FROM salaries;
		{
			mgorm.Avg(db, "salary").
				From("salaries").(*mgorm.SelectStmt),
			&avg,
		},
	}

	for i, testCase := range testCases {
		model := new(float64)
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
