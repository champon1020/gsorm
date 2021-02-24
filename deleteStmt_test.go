package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("sample").(*mgorm.DeleteStmt),
			`DELETE FROM sample`,
		},
		{
			mgorm.Delete(nil).From("sample").
				Where("id = ?", 10000).(*mgorm.DeleteStmt),
			`DELETE FROM sample WHERE id = 10000`,
		},
		{
			mgorm.Delete(nil).From("sample").
				Where("id = ?", 10000).
				And("name = ? OR name = ?", "Taro", "Jiro").(*mgorm.DeleteStmt),
			`DELETE FROM sample WHERE id = 10000 AND (name = "Taro" OR name = "Jiro")`,
		},
		{
			mgorm.Delete(nil).From("sample").
				Where("id = ?", 10000).
				Or("name = ? AND nickname = ?", "Taro", "TaroTaro").(*mgorm.DeleteStmt),
			`DELETE FROM sample WHERE id = 10000 OR (name = "Taro" AND nickname = "TaroTaro")`,
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

func TestDeleteStmt_ProcessSQL_Fail(t *testing.T) {
	{
		expectedErr := errors.New(
			"clause.Join is not supported for DELETE statement", errors.InvalidSyntaxError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Delete(nil).(*mgorm.DeleteStmt)
		s.ExportedSetCalled(&clause.Join{})

		// Actual process.
		var sql internal.SQL
		err := mgorm.DeleteStmtProcessSQL(s, &sql)

		// Validate error.
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
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
