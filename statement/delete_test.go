package statement_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
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
				s := mgorm.Delete(nil).(*statement.DeleteStmt)
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
			mgorm.Delete(nil).From("table").(*statement.DeleteStmt),
			mgorm.Delete(nil).From("table").Where("column1 = ?", 10).(*statement.DeleteStmt),
			statement.ErrInvalidValue,
		},
		{
			mgorm.Delete(nil).From("table").Where("column1 = ?", 10).(*statement.DeleteStmt),
			mgorm.Delete(nil).From("table").Where("column1 = ?", 100).(*statement.DeleteStmt),
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
			mgorm.Delete(nil).
				RawClause("RAW").
				From("employees").(*statement.DeleteStmt),
			`DELETE RAW FROM employees`,
		},
		{
			mgorm.Delete(nil).
				From("employees").
				RawClause("RAW").
				Where("emp_no = ?", 10000).(*statement.DeleteStmt),
			`DELETE FROM employees RAW WHERE emp_no = 10000`,
		},
		{
			mgorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				And("first_name = ?", "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW AND (first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				Or("emp_no = ?", 20000).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW OR (emp_no = 20000)`,
		},
		{
			mgorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				And("first_name = ?", "Taro").
				RawClause("RAW").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 AND (first_name = 'Taro') RAW`,
		},
		{
			mgorm.Delete(nil).
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
