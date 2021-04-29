package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
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
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Delete(nil).(*mgorm.DeleteStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.DeleteStmtBuildSQL(s, &sql)
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

func TestInsertStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").(*mgorm.InsertStmt)
				s.ExportedSetCalled(&clause.Set{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
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
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").Model(1000).(*mgorm.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := mgorm.Insert(nil, "table", "column").Model(&model).(*mgorm.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
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

func TestSelectStmt_BuildQuerySQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Select(nil, "").(*mgorm.SelectStmt)
				s.ExportedSetCalled(&clause.Values{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.SelectStmtBuildSQL(s, &sql)
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

func TestUpdateStmt_String(t *testing.T) {
	type Model struct {
		ID   int
		Name string `mgorm:"first_name"`
	}

	model1 := Model{ID: 10000, Name: "Taro"}
	model2 := map[string]interface{}{"id": 10000, "first_name": "Taro"}

	testCases := []struct {
		Stmt     *mgorm.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "sample").
				Set("id", 10000).
				Set("first_name", "Taro").
				Where("id = ?", 20000).
				And("first_name = ? OR first_name = ?", "Jiro", "Hanako").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000 ` +
				`AND (first_name = 'Jiro' OR first_name = 'Hanako')`,
		},
		{
			mgorm.Update(nil, "sample").
				Set("id", 10000).
				Set("first_name", "Taro").
				Where("id = ?", 20000).
				Or("first_name = ? AND last_name = ?", "Jiro", "Sato").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000 ` +
				`OR (first_name = 'Jiro' AND last_name = 'Sato')`,
		},
		{
			mgorm.Update(nil, "sample").
				Model(&model1, "id", "first_name").
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000`,
		},
		{
			mgorm.Update(nil, "sample").
				Model(&model2, "id", "first_name").
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000`,
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

func TestUpdateStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Update(nil, "table").(*mgorm.UpdateStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
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
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := map[string]interface{}{
					"id":   1000,
					"name": "Taro",
				}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*mgorm.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := []int{1000}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*mgorm.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
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
