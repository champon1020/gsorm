package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
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

func TestUpdateStmt_Set_Fail(t *testing.T) {
	{
		expectedErr := errors.New("Command is nil", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := new(mgorm.UpdateStmt)

		// Actual process.
		s.Set("")

		// Validate error.
		errs := s.ExportedGetErrors()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		expectedErr := errors.New(
			"Length is different between columns and values", errors.InvalidValueError).(*errors.Error)

		// Actual process.
		s := mgorm.Update(nil, "sample", "id").Set(10, "Taro").(*mgorm.UpdateStmt)

		// Validate error.
		errs := s.ExportedGetErrors()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}
