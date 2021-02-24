package dockertest_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/google/go-cmp/cmp"
)

func TestSelectWhere(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.SelectStmt
		Result *[]Employee
	}{
		// SELECT emp_no FROM employees WHERE emp_no = 10001;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Where("emp_no = ?", 10001).(*mgorm.SelectStmt),
			&[]Employee{
				{EmpNo: 10001},
			},
		},

		// SELECT emp_no, first_name, last_name FROM employees
		// WHERE emp_no <= 10005 AND (first_name = "Georgi" OR last_name = "Bamford");
		{
			mgorm.Select(db, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10005).
				And("first_name = ? OR last_name = ?", "Georgi", "Bamford").(*mgorm.SelectStmt),
			&[]Employee{
				{
					EmpNo:     10001,
					FirstName: "Georgi",
					LastName:  "Facello",
				},
				{
					EmpNo:     10003,
					FirstName: "Parto",
					LastName:  "Bamford",
				},
			},
		},

		// SELECT emp_no, first_name, last_name FROM employees
		// WHERE emp_no <= 10002 OR (first_name = "Saniya" AND last_name = "Kalloufi");
		{
			mgorm.Select(db, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10002).
				Or("first_name = ? AND last_name = ?", "Saniya", "Kalloufi").(*mgorm.SelectStmt),
			&[]Employee{
				{
					EmpNo:     10001,
					FirstName: "Georgi",
					LastName:  "Facello",
				},
				{
					EmpNo:     10002,
					FirstName: "Bezalel",
					LastName:  "Simmel",
				},
				{
					EmpNo:     10008,
					FirstName: "Saniya",
					LastName:  "Kalloufi",
				},
			},
		},

		// SELECT emp_no FROM employees
		// WHERE emp_no IN (SELECT DISTINCT emp_no FROM salaries WHERE salary < 60000);
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Where("emp_no IN ?",
					mgorm.Select(nil, "DISTINCT emp_no").
						From("salaries").
						Where("salary < ?", 60000),
				).(*mgorm.SelectStmt),
			&[]Employee{
				{EmpNo: 10004},
				{EmpNo: 10007},
			},
		},

		// SELECT emp_no FROM employees WHERE emp_no BETWEEN 10002 AND 10004;
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Where("emp_no BETWEEN ? AND ?", 10002, 10004).(*mgorm.SelectStmt),
			&[]Employee{
				{EmpNo: 10002},
				{EmpNo: 10003},
				{EmpNo: 10004},
			},
		},

		// SELECT first_name FROM employees WHERE first_name LIKE "S%";
		{
			mgorm.Select(db, "first_name").
				From("employees").
				Where("first_name LIKE ?", "S%").(*mgorm.SelectStmt),
			&[]Employee{
				{FirstName: "Saniya"},
				{FirstName: "Sumant"},
			},
		},

		// SELECT emp_no FROM employees WHERE emp_no IN (10002, 10004);
		{
			mgorm.Select(db, "emp_no").
				From("employees").
				Where("emp_no IN (?, ?)", 10002, 10004).(*mgorm.SelectStmt),
			&[]Employee{
				{EmpNo: 10002},
				{EmpNo: 10004},
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
