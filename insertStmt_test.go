package mgorm_test

/*
func TestInsertStmt_ProcessSQL_Fail(t *testing.T) {
	{
		expectedErr := errors.New(
			"Type clause.Join is not supported for INSERT", errors.InvalidTypeError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Insert(nil, "").(*mgorm.InsertStmt).Join("").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.InsertStmtProcessSQL(s)

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
*/
