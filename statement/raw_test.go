package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement"
	"github.com/stretchr/testify/assert"
)

func TestRawStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.RawStmt
		Expected string
	}{
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees").(*statement.RawStmt),
			`SELECT * FROM employees`,
		},
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no = ?", 1001).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees WHERE first_name = ?", "Taro").(*statement.RawStmt),
			`SELECT * FROM employees WHERE first_name = 'Taro'`,
		},
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees WHERE birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.RawStmt),
			`SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)", []int{1001, 1002}).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.RawStmt),
			`SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
		{
			mgorm.RawStmt(nil, "DELETE FROM employees").(*statement.RawStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.RawStmt(nil, "ALTER TABLE employees DROP PRIMARY KEY PK_emp_no").(*statement.RawStmt),
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
