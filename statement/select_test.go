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

func TestSelectStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).
				RawClause("RAW").
				From("table").(*statement.SelectStmt),
			`SELECT * RAW FROM table`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				Join("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`INNER JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				LeftJoin("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`LEFT JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				RightJoin("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`RIGHT JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Join("table2").RawClause("RAW").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 RAW ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Join("table2").On("table.column = table2.column").
				RawClause("RAW").
				Where("id = ?", 10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 ON table.column = table2.column ` +
				`RAW WHERE id = 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				And("name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW AND (name = 'Taro')`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				Or("name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW OR (name = 'Taro')`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				And("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 AND (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				Or("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 OR (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				GroupBy("column").
				RawClause("RAW").
				Having("SUM(id) = ?", 10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`GROUP BY column ` +
				`RAW HAVING SUM(id) = 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				Union(mgorm.Select(nil).From("table2")).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION (SELECT * FROM table2)`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				UnionAll(mgorm.Select(nil).From("table2")).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION ALL (SELECT * FROM table2)`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Union(mgorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`UNION (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			mgorm.Select(nil).
				From("table").
				UnionAll(mgorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`UNION ALL (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			mgorm.Select(nil).
				From("table").
				OrderBy("id").
				RawClause("RAW").
				Limit(10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`ORDER BY id ` +
				`RAW LIMIT 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Limit(10).
				RawClause("RAW").
				Offset(5).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 RAW OFFSET 5`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Limit(10).
				Offset(5).
				RawClause("RAW").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 OFFSET 5 RAW`,
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
