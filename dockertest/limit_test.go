package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestSelectLimit(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]Employee
	}{
		// SELECT emp_no FROM employees LIMIT 5;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Limit(5).(*mgorm.Stmt),
			&[]Employee{
				{EmpNo: 10001},
				{EmpNo: 10002},
				{EmpNo: 10003},
				{EmpNo: 10004},
				{EmpNo: 10005},
			},
		},

		// SELECT emp_no FROM employees LIMIT 5 OFFSET 3;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Limit(5).
				Offset(3).(*mgorm.Stmt),
			&[]Employee{
				{EmpNo: 10004},
				{EmpNo: 10005},
				{EmpNo: 10006},
				{EmpNo: 10007},
				{EmpNo: 10008},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Employee)
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
