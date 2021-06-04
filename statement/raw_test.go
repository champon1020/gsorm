package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestRawStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.RawStmt
		Expected string
	}{
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees").(*statement.RawStmt),
			`SELECT * FROM employees`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no = ?", 1001).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE first_name = ?", "Taro").(*statement.RawStmt),
			`SELECT * FROM employees WHERE first_name = 'Taro'`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE birth_date = ?",
				time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.RawStmt),
			`SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)", []int{1001, 1002}).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)",
				gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
		{
			gsorm.RawStmt(nil, "DELETE FROM employees").(*statement.RawStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.RawStmt(nil, "ALTER TABLE employees DROP PRIMARY KEY PK_emp_no").(*statement.RawStmt),
			`ALTER TABLE employees DROP PRIMARY KEY PK_emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestRawStmt_QueryWithMock(t *testing.T) {
	type Employee struct {
		EmpNo     int
		FirstName string
	}
	model := []Employee{}
	expectedReturn := []Employee{{1001, "Taro"}, {1002, "Jiro"}}

	mock := gsorm.OpenMock("mysql")
	mock.ExpectWithReturn(gsorm.RawStmt(nil, "SELECT emp_no, first_name FROM employees"), expectedReturn)

	if err := gsorm.RawStmt(mock, "SELECT emp_no, first_name FROM employees").Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if diff := cmp.Diff(expectedReturn, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRawStmt_ExecWithMock(t *testing.T) {
	mock := gsorm.OpenMock("mysql")
	mock.Expect(gsorm.RawStmt(mock, `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`))

	if err := gsorm.RawStmt(mock, `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`).Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}
}
