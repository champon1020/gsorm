package database_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/statement"
	"github.com/morikuni/failure"
)

func TestCompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.SelectStmt
		ActualStmt    *statement.SelectStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Select(nil, "column1").From("table").(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			database.ErrInvalidMockExpectation,
		},
		{
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 100).(*statement.SelectStmt),
			database.ErrInvalidMockExpectation,
		},
	}

	for _, testCase := range testCases {
		err := database.CompareStmts(testCase.ExpectedStmt, testCase.ActualStmt)

		// Validate if the expected error was occurred.
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}
