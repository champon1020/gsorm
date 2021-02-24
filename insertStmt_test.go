package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/stretchr/testify/assert"
)

func TestInsertStmt_String(t *testing.T) {
	type Model struct {
		ID        int
		FirstName string `mgorm:"name"`
	}

	model1 := Model{ID: 10000, FirstName: "Taro"}
	model2 := []Model{{ID: 10000, FirstName: "Taro"}, {ID: 10001, FirstName: "Jiro"}}
	model3 := []int{10000, 10001}
	model4 := map[string]interface{}{
		"id":   10000,
		"name": "Taro",
	}

	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "sample", "id", "name").Values(10000, "Taro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").
				Values(10002, "Saburo").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro"), (10002, "Saburo")`,
		},
		// Test for (*InsertStmt).Model
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro"), (10001, "Jiro")`,
		},
		{
			mgorm.Insert(nil, "sample", "id").Model(&model3).(*mgorm.InsertStmt),
			`INSERT INTO sample (id) VALUES (10000), (10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, "Taro")`,
		},
		// Test for mapping.
		{
			mgorm.Insert(nil, "sample", "first_name AS name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (first_name AS name, id) VALUES ("Taro", 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000), ("Jiro", 10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ("Taro", 10000)`,
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

func TestInsertStmt_ProcessSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedErr *errors.Error
		Process     func() error
	}{
		{
			errors.New("clause.Set is not supported for INSERT statement", errors.InvalidSyntaxError).(*errors.Error),
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").(*mgorm.InsertStmt)
				s.ExportedSetCalled(&clause.Set{})

				// Actual process.
				var sql internal.SQL
				err := mgorm.InsertStmtProcessSQL(s, &sql)
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

func TestInsertStmt_ProcessSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedErr *errors.Error
		Process     func() error
	}{
		{
			errors.New("Model must be pointer", errors.InvalidValueError).(*errors.Error),
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").Model(1000).(*mgorm.InsertStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.InsertStmtProcessSQL(s, &sql)
				return err
			},
		},
		{
			errors.New("Column names must be included in one of map keys", errors.InvalidSyntaxError).(*errors.Error),
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := mgorm.Insert(nil, "table", "column").Model(&model).(*mgorm.InsertStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.InsertStmtProcessSQL(s, &sql)
				return err
			},
		},
		{
			errors.New("Type *int is not supported for (*InsertStmt).Model", errors.InvalidTypeError).(*errors.Error),
			func() error {
				// Prepare for test.
				model := 10000
				s := mgorm.Insert(nil, "table", "column").Model(&model).(*mgorm.InsertStmt)

				// Actual process.
				var sql internal.SQL
				err := mgorm.InsertStmtProcessSQL(s, &sql)
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
