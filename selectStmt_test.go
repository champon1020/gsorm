package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
)

func TestStmt_ProcessQuerySQL_Fail(t *testing.T) {
	{
		expectedErr := errors.New("Command must be SELECT", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Update(nil, "column1", "column2").Set(10, "str").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.SelectStmtProcessSQL(s)

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
	{
		expectedErr := errors.New(
			`Type clause.Values is not supported for SELECT`, errors.InvalidTypeError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Select(nil, "").(*mgorm.Stmt).Values("").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.SelectStmtProcessSQL(s)

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
