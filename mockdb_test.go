package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/google/go-cmp/cmp"
)

func TestMock_Expectation(t *testing.T) {
	expectedReturn := []int{10, 20, 30}

	// Test phase.
	mock := mgorm.NewMock()
	mock.Expect(mgorm.Insert(nil, "table", "column1", "column2").Values(10, "str"))
	mock.Expect(mgorm.Select(nil, "column1").From("table")).Return(expectedReturn)

	// Actual process.
	if err := mgorm.Insert(mock, "table", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	model := new([]int)
	if err := mgorm.Select(mock, "column1").From("table").Query(model); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Test phase.
	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate model values.
	if diff := cmp.Diff(*model, expectedReturn); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestMock_TransactionExpectation(t *testing.T) {
	expectedReturn1 := []int{10, 20, 30}
	expectedReturn2 := []string{"hello", "world", "!"}

	// Test phase.
	mock := mgorm.NewMock()
	mocktx1 := mock.ExpectBegin()
	mocktx2 := mock.ExpectBegin()

	mocktx1.Expect(mgorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mocktx1.Expect(mgorm.Select(nil, "column1").From("table1")).Return(expectedReturn1)
	mocktx1.ExpectCommit()
	mocktx2.Expect(mgorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))
	mocktx2.Expect(mgorm.Select(nil, "column2").From("table2")).Return(expectedReturn2)
	mocktx2.ExpectRollback()

	// Actual process.
	tx1, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	tx2, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	if err := mgorm.Insert(tx1, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	model1 := new([]int)
	if err := mgorm.Select(tx1, "column1").From("table1").Query(model1); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	tx1.Commit()

	if err := mgorm.Insert(tx2, "table2", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	model2 := new([]string)
	if err := mgorm.Select(tx2, "column2").From("table2").Query(model2); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}
	tx2.Rollback()

	// Test phase.
	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate model values.
	if diff := cmp.Diff(*model1, expectedReturn1); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
	if diff := cmp.Diff(*model2, expectedReturn2); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestMockDB_Begin_Fail(t *testing.T) {
	expectedErr1 := errors.New(
		"mgorm.(*MockDB).Begin was executed but not expected", errors.MockError).(*errors.Error)

	expectedErr2 := errors.New(
		`mgorm.(*MockDB).Begin was executed but SELECT("column1").FROM("table") is expected`, errors.MockError).(*errors.Error)

	{
		// Test phase.
		mock := new(mgorm.MockDB)

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr1) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr1.Error(), expectedErr1.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		// Test phase.
		mock := new(mgorm.MockDB)
		mock.Expect(mgorm.Select(nil, "column1").From("table"))
		_ = mock.ExpectBegin()

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr2) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr2.Error(), expectedErr2.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestMockDB_Complete_Fail(t *testing.T) {
	expectedErr := errors.New(`No query was executed, but `+
		`INSERT INTO("table2", "column1", "column2").VALUES(10, 'str') `+
		`is expected`, errors.MockError).(*errors.Error)

	// Test phase.
	mock := mgorm.NewMock()
	mock.Expect(mgorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mock.Expect(mgorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	if err := mgorm.Insert(mock, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Test phase.
	err := mock.Complete()
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

func TestMockDB_Complete_Transaction_Fail(t *testing.T) {
	expectedErr := errors.New(`No query was executed, but `+
		`INSERT INTO("table2", "column1", "column2").VALUES(10, 'str') `+
		`is expected`, errors.MockError).(*errors.Error)

	// Test phase.
	mock := mgorm.NewMock()
	mocktx := mock.ExpectBegin()
	mocktx.Expect(mgorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mocktx.Expect(mgorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	tx, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occured: %v", err)
		return
	}
	if err := mgorm.Insert(tx, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Test phase.
	err = mock.Complete()
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

func TestMockDB_CompareWith(t *testing.T) {
	expectedErr1 := errors.New(
		`SELECT("column1").FROM("table") was executed but not expected`, errors.MockError).(*errors.Error)

	expectedErr2 := errors.New(
		`SELECT("column1").FROM("table") was executed `+
			`but mgorm.(*MockDB).Begin is expected`, errors.MockError).(*errors.Error)

	{
		// Test phase.
		mock := mgorm.NewMock()

		// Actual process.
		model := new([]int)
		err := mgorm.Select(mock, "column1").From("table").Query(model)
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr1) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr1.Error(), expectedErr1.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		err := mgorm.Select(mock, "column1").From("table").Query(model)
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr2) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr2.Error(), expectedErr2.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestMockTx_Commit_Fail(t *testing.T) {
	expectedErr := errors.New(
		"mgorm.(*MockTx).Commit was executed but not expected", errors.MockError).(*errors.Error)

	{
		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %v", err)
			return
		}
		err = tx.Commit()
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
		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectRollback()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %v", err)
			return
		}
		err = tx.Commit()
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

func TestMockTx_Rollback_Fail(t *testing.T) {
	expectedErr := errors.New(
		"mgorm.(*MockTx).Rollback was executed but not expected", errors.MockError).(*errors.Error)

	{
		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %v", err)
			return
		}
		err = tx.Rollback()
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
		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectCommit()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %v", err)
			return
		}
		err = tx.Rollback()
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

func TestMockTx_CompareWith(t *testing.T) {
	expectedErr1 := errors.New(
		`SELECT("column1").FROM("table") was executed but not expected`, errors.MockError).(*errors.Error)

	expectedErr2 := errors.New(
		`SELECT("column1").FROM("table") was executed `+
			`but mgorm.(*MockTx).Commit is expected`, errors.MockError).(*errors.Error)

	{
		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			return
		}
		err = mgorm.Select(tx, "column1").From("table").Query(model)
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr1) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr1.Error(), expectedErr1.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectCommit()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			return
		}
		err = mgorm.Select(tx, "column1").From("table").Query(model)
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr2) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr2.Error(), expectedErr2.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestCompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *mgorm.SelectStmt
		ActualStmt    *mgorm.SelectStmt
		ExpectedError error
	}{
		{
			mgorm.Select(nil, "column1").From("table").(*mgorm.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*mgorm.SelectStmt),
			errors.New(`SELECT("column1").FROM("table").WHERE("column1 = ?", 10) was executed `+
				`but SELECT("column1").FROM("table") is expected`, errors.MockError),
		},
		{
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*mgorm.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 100).(*mgorm.SelectStmt),
			errors.New(`SELECT("column1").FROM("table").WHERE("column1 = ?", 100) was executed `+
				`but SELECT("column1").FROM("table").WHERE("column1 = ?", 10) is expected`, errors.MockError),
		},
	}

	for _, testCase := range testCases {
		err := mgorm.CompareStmts(testCase.ExpectedStmt, testCase.ActualStmt)
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		expectedErr := testCase.ExpectedError.(*errors.Error)
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)

		}
	}
}
