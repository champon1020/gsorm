package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/stretchr/testify/assert"
)

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
			mgorm.Update(nil, "sample", "id", "first_name").
				Set(10000, "Taro").
				Where("id = ?", 20000).
				And("first_name = ? OR first_name = ?", "Jiro", "Hanako").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = "Taro" ` +
				`WHERE id = 20000 ` +
				`AND (first_name = "Jiro" OR first_name = "Hanako")`,
		},
		{
			mgorm.Update(nil, "sample", "id", "first_name").
				Set(10000, "Taro").
				Where("id = ?", 20000).
				Or("first_name = ? AND last_name = ?", "Jiro", "Sato").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = "Taro" ` +
				`WHERE id = 20000 ` +
				`OR (first_name = "Jiro" AND last_name = "Sato")`,
		},
		{
			mgorm.Update(nil, "sample", "id", "first_name").
				Model(&model1).
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = "Taro" ` +
				`WHERE id = 20000`,
		},
		{
			mgorm.Update(nil, "sample", "id").
				Model(10000).
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000 ` +
				`WHERE id = 20000`,
		},
		{
			mgorm.Update(nil, "sample", "id", "first_name").
				Model(&model2).
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = "Taro" ` +
				`WHERE id = 20000`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestUpdateStmt_ProcessSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedErr *errors.Error
		Process     func() error
	}{
		{
			errors.New("clause.Join is not supported for UPDATE statement", errors.InvalidTypeError).(*errors.Error),
			func() error {
				// Prepare for test.
				s := mgorm.Update(nil, "", "").(*mgorm.UpdateStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual process.
				var sql internal.SQL
				err := mgorm.UpdateStmtProcessSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Process()
		if err == nil {
			t.Errorf("Error was not occurred")
			continue
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		if !actualErr.Is(testCase.ExpectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", testCase.ExpectedErr.Error(), testCase.ExpectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestUpdateStmt_ProcessSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedErr *errors.Error
		Process     func() error
	}{
		{
			errors.New(
				"If you set variable to Model, number of columns must be 1, not 2",
				errors.InvalidSyntaxError).(*errors.Error),
			func() error {
				// Prepare for test.
				s := mgorm.Update(nil, "table", "column1", "column2").Model(1000).(*mgorm.UpdateStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.UpdateStmtProcessSQL(s, &sql)
				return err
			},
		},
		{
			errors.New("If model is not variable, model must be pointer", errors.InvalidValueError).(*errors.Error),
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := mgorm.Update(nil, "table", "column").Model(model).(*mgorm.UpdateStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.UpdateStmtProcessSQL(s, &sql)
				return err
			},
		},
		{
			errors.New("Column names must be included in one of map keys", errors.InvalidSyntaxError).(*errors.Error),
			func() error {
				// Prepare for test.
				model := map[string]interface{}{
					"id":   1000,
					"name": "Taro",
				}
				s := mgorm.Update(nil, "sample", "id", "first_name").Model(&model).(*mgorm.UpdateStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.UpdateStmtProcessSQL(s, &sql)
				return err
			},
		},
		{
			errors.New("Type *[]int is not supported for (*UpdateStmt).Model", errors.InvalidTypeError).(*errors.Error),
			func() error {
				// Prepare for test.
				model := []int{1000}
				s := mgorm.Update(nil, "sample", "id", "first_name").Model(&model).(*mgorm.UpdateStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.UpdateStmtProcessSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Process()
		if err == nil {
			t.Errorf("Error was not occurred")
			continue
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		if !actualErr.Is(testCase.ExpectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", testCase.ExpectedErr.Error(), testCase.ExpectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestUpdateStmt_Set_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedErr *errors.Error
		Process     func() []error
	}{
		{
			errors.New("(*UpdateStmt).cmd is nil", errors.InvalidValueError).(*errors.Error),
			func() []error {
				// Prepare for test.
				s := new(mgorm.UpdateStmt)

				// Actual process.
				s.Set("")
				return s.ExportedGetErrors()
			},
		},
		{
			errors.New("Number of values is not equal to that of columns", errors.InvalidValueError).(*errors.Error),
			func() []error {
				// Actual process.
				s := mgorm.Update(nil, "sample", "id").Set(10, "Taro").(*mgorm.UpdateStmt)
				return s.ExportedGetErrors()
			},
		},
	}

	for _, testCase := range testCases {
		errs := testCase.Process()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(testCase.ExpectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", testCase.ExpectedErr.Error(), testCase.ExpectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}
