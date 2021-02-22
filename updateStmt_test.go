package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
)

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
}
