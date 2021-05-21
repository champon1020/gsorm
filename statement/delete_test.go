package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/statement"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_BuildSQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := gsorm.Delete(nil).(*statement.DeleteStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := statement.DeleteStmtBuildSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Build()
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestDeleteStmt_CompareStmts_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.DeleteStmt
		ActualStmt    *statement.DeleteStmt
		ExpectedError failure.StringCode
	}{
		{
			gsorm.Delete(nil).From("table").(*statement.DeleteStmt),
			gsorm.Delete(nil).From("table").Where("column1 = ?", 10).(*statement.DeleteStmt),
			statement.ErrInvalidValue,
		},
		{
			gsorm.Delete(nil).From("table").Where("column1 = ?", 10).(*statement.DeleteStmt),
			gsorm.Delete(nil).From("table").Where("column1 = ?", 100).(*statement.DeleteStmt),
			statement.ErrInvalidValue,
		},
	}

	for _, testCase := range testCases {
		err := testCase.ExpectedStmt.CompareWith(testCase.ActualStmt)

		// Validate if the expected error was occurred.
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestDeleteStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).
				RawClause("RAW").
				From("employees").(*statement.DeleteStmt),
			`DELETE RAW FROM employees`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				RawClause("RAW").
				Where("emp_no = ?", 10000).(*statement.DeleteStmt),
			`DELETE FROM employees RAW WHERE emp_no = 10000`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				And("first_name = ?", "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW AND (first_name = 'Taro')`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				Or("emp_no = ?", 20000).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW OR (emp_no = 20000)`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				And("first_name = ?", "Taro").
				RawClause("RAW").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 AND (first_name = 'Taro') RAW`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				Or("emp_no = ?", 20000).
				RawClause("RAW").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 OR (emp_no = 20000) RAW`,
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

func TestDeleteStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").(*statement.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.Delete(nil).From("employees").(*statement.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.Delete(nil).From("employees AS e").(*statement.DeleteStmt),
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
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = 1001").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("first_name = ?", "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE first_name = 'Taro'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
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
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
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
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
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
