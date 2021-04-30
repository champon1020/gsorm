package statement_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
)

func TestSelectStmt_BuildQuerySQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Select(nil, "").(*statement.SelectStmt)
				s.ExportedSetCalled(&clause.Values{})

				// Actual build.
				var sql internal.SQL
				err := statement.SelectStmtBuildSQL(s, &sql)
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

func TestSelectStmt_CompareStmts_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.SelectStmt
		ActualStmt    *statement.SelectStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Select(nil, "column1").From("table").(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			statement.ErrInvalidValue,
		},
		{
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 100).(*statement.SelectStmt),
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
