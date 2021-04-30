package statement_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
)

func TestUpdateStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Update(nil, "table").(*statement.UpdateStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
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

func TestUpdateStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := map[string]interface{}{
					"id":   1000,
					"name": "Taro",
				}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*statement.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := []int{1000}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*statement.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
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

func TestUpdateStmt_CompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.UpdateStmt
		ActualStmt    *statement.UpdateStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Update(nil, "table").Set("col1", 10).(*statement.UpdateStmt),
			mgorm.Update(nil, "table").Set("col1", 10).Set("col2", 100).(*statement.UpdateStmt),
			statement.ErrInvalidValue,
		},
		{
			mgorm.Update(nil, "table").Set("col1", 10).(*statement.UpdateStmt),
			mgorm.Update(nil, "table").Set("col1", 100).(*statement.UpdateStmt),
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
