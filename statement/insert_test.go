package statement_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
)

func TestInsertStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").(*statement.InsertStmt)
				s.ExportedSetCalled(&clause.Set{})

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
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

func TestInsertStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").Model(1000).(*statement.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := mgorm.Insert(nil, "table", "column").Model(&model).(*statement.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
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

func TestInsertStmt_CompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.InsertStmt
		ActualStmt    *statement.InsertStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Insert(nil, "table").Values(10).(*statement.InsertStmt),
			mgorm.Insert(nil, "table").Values(10).Values(100).(*statement.InsertStmt),
			statement.ErrInvalidValue,
		},
		{
			mgorm.Insert(nil, "table").Values(10).(*statement.InsertStmt),
			mgorm.Insert(nil, "table").Values(10, 100).(*statement.InsertStmt),
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
