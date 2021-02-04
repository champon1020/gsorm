package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestMaxMin(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *Employee
	}{
		// SELECT MIN(emp_no) FROM employees;
		{
			mgorm.Min(db, "emp_no").
				From("employees").(*mgorm.Stmt),
			&Employee{
				EmpNo: 10001,
			},
		},

		// SELECT MAX(emp_no) FROM employees;
		{
			mgorm.Max(db, "emp_no").
				From("employees").(*mgorm.Stmt),
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
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestCountSum(t *testing.T) {
	var (
		cnt = 10
		sum = 579290
	)

	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *int
	}{
		{
			mgorm.Count(db, "emp_no").
				From("salaries").(*mgorm.Stmt),
			&cnt,
		},
		{
			mgorm.Sum(db, "salary").
				From("salaries").(*mgorm.Stmt),
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
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestAvg(t *testing.T) {
	avg := 57929.0

	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *float64
	}{
		{
			mgorm.Avg(db, "salary").
				From("salaries").(*mgorm.Stmt),
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
			internal.PrintTestDiff(t, diff)
		}
	}
}
