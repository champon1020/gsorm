package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").(*mgorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees").(*mgorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees AS e").(*mgorm.DeleteStmt),
			`DELETE FROM employees AS e`,
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

func TestDeleteStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").Where("emp_no = 1001").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Delete(nil).From("employees").Where("emp_no = ?", 1001).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
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

func TestDeleteStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ?", "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR last_name = ?", "Taro", "Sato").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR last_name = 'Sato')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ?", "Taro").
				And("last_name = ?", "Sato").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro') AND (last_name = 'Sato')`,
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

func TestDeleteStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
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
