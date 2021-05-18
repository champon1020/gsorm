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

func TestInsertStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "table").
				RawClause("RAW").
				Values("value").(*statement.InsertStmt),
			`INSERT INTO table RAW VALUES ('value')`,
		},
		{
			mgorm.Insert(nil, "table").
				Values("value1").
				RawClause("RAW").
				Values("value2").(*statement.InsertStmt),
			`INSERT INTO table VALUES ('value1') RAW, ('value2')`,
		},
		{
			mgorm.Insert(nil, "table").
				Values("value").
				RawClause("RAW").(*statement.InsertStmt),
			`INSERT INTO table VALUES ('value') RAW`,
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
