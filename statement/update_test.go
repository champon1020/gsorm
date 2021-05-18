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

func TestUpdateStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "table").
				RawClause("RAW").
				Set("column", "value").(*statement.UpdateStmt),
			`UPDATE table RAW SET column = 'value'`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Set("column2", "value2").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' RAW, column2 = 'value2'`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Where("column = ?", 10).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`RAW WHERE column = 10`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				And("column = ?", 100).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW AND (column = 100)`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				Or("column = ?", 100).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW OR (column = 100)`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				And("column = ?", 100).
				RawClause("RAW").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`AND (column = 100) RAW`,
		},
		{
			mgorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				Or("column = ?", 100).
				RawClause("RAW").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`OR (column = 100) RAW`,
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
