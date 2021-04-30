package statement_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
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
