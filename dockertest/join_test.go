package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/google/go-cmp/cmp"
)

type EmployeeWithTitle struct {
	EmpNo int
	Title string
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]EmployeeWithTitle
	}{
		// SELECT e.emp_no, t.title FROM employees AS e
		// INNER JOIN titles AS t ON e.emp_no = t.emp_no
		// LIMIT 5;
		{
			mgorm.Select(db, "e.emp_no", "t.title").
				From("employees AS e").
				Join("titles AS t").
				On("e.emp_no = t.emp_no").
				Limit(5).(*mgorm.Stmt),
			&[]EmployeeWithTitle{
				{EmpNo: 10001, Title: "Senior Engineer"},
				{EmpNo: 10002, Title: "Staff"},
				{EmpNo: 10003, Title: "Senior Engineer"},
				{EmpNo: 10004, Title: "Engineer"},
				{EmpNo: 10004, Title: "Senior Engineer"},
			},
		},

		// SELECT e.emp_no, t.title FROM employees AS e
		// LEFT JOIN titles AS t ON e.emp_no = t.emp_no
		// ORDER BY e.emp_no DESC LIMIT 5;
		{
			mgorm.Select(db, "e.emp_no", "t.title").
				From("employees AS e").
				LeftJoin("titles AS t").
				On("e.emp_no = t.emp_no").
				OrderBy("e.emp_no DESC", "title").
				Limit(5).(*mgorm.Stmt),
			&[]EmployeeWithTitle{
				{EmpNo: 10010},
				{EmpNo: 10009},
				{EmpNo: 10008},
				{EmpNo: 10007, Title: "Senior Staff"},
				{EmpNo: 10007, Title: "Staff"},
			},
		},

		// SELECT e.emp_no, t.title FROM employees AS e
		// RIGHT JOIN titles AS t ON e.emp_no = t.emp_no
		// ORDER BY e.emp_no DESC LIMIT 5;
		{
			mgorm.Select(db, "t.title", "e.emp_no").
				From("titles AS t").
				RightJoin("employees AS e").
				On("t.emp_no = e.emp_no").
				OrderBy("e.emp_no DESC", "title").
				Limit(5).(*mgorm.Stmt),
			&[]EmployeeWithTitle{
				{EmpNo: 10010},
				{EmpNo: 10009},
				{EmpNo: 10008},
				{EmpNo: 10007, Title: "Senior Staff"},
				{EmpNo: 10007, Title: "Staff"},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]EmployeeWithTitle)
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
