package e2e_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement"
	"github.com/google/go-cmp/cmp"
)

func TestSelectLimit(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]Employee
	}{
		// SELECT emp_no FROM employees LIMIT 5;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Limit(5).(*statement.SelectStmt),
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
				Offset(3).(*statement.SelectStmt),
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
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
