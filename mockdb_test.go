package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
)

func TestMock_Expectation(t *testing.T) {
	expectedReturn := []int{10, 20, 30}

	// Test phase.
	mock := mgorm.NewMock()
	mock.Expect(mgorm.Insert(nil, "table", "column1", "column2").Values(10, "str"))
	mock.Expect(mgorm.Select(nil, "column1").From("table")).Return(expectedReturn)

	// Actual process.
	if err := mgorm.Insert(mock, "table", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	model := new([]int)
	if err := mgorm.Select(mock, "column1").From("table").Query(model); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	// Test phase.
	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
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
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	tx2, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := mgorm.Insert(tx1, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	model1 := new([]int)
	if err := mgorm.Select(tx1, "column1").From("table1").Query(model1); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := tx1.Commit(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := mgorm.Insert(tx2, "table2", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	model2 := new([]string)
	if err := mgorm.Select(tx2, "column2").From("table2").Query(model2); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := tx2.Rollback(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	// Test phase.
	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
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
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := new(mgorm.MockDB)

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := new(mgorm.MockDB)
		mock.Expect(mgorm.Select(nil, "column1").From("table"))
		_ = mock.ExpectBegin()

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockDB_Complete_Fail(t *testing.T) {
	expectedErr := mgorm.ErrInvalidMockExpectation

	// Test phase.
	mock := mgorm.NewMock()
	mock.Expect(mgorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mock.Expect(mgorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	if err := mgorm.Insert(mock, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	// Test phase.
	err := mock.Complete()

	// Validate if the expected error was occurred.
	if !failure.Is(err, expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedErr)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestMockDB_Complete_Transaction_Fail(t *testing.T) {
	expectedErr := mgorm.ErrInvalidMockExpectation

	// Test phase.
	mock := mgorm.NewMock()
	mocktx := mock.ExpectBegin()
	mocktx.Expect(mgorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mocktx.Expect(mgorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	tx, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occured: %+v", err)
		return
	}
	if err = mgorm.Insert(tx, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	// Test phase.
	err = mock.Complete()

	// Validate if the expected error was occurred.
	if !failure.Is(err, expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedErr)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestMockDB_CompareWith(t *testing.T) {
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()

		// Actual process.
		model := new([]int)
		err := mgorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		err := mgorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockTx_Commit_Fail(t *testing.T) {
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Commit()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectRollback()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Commit()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockTx_Rollback_Fail(t *testing.T) {
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Rollback()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectCommit()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Rollback()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockTx_CompareWith(t *testing.T) {
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %+v", err)
			return
		}
		err = mgorm.Select(tx, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := mgorm.ErrInvalidMockExpectation

		// Test phase.
		mock := mgorm.NewMock()
		mocktx := mock.ExpectBegin()
		mocktx.ExpectCommit()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %+v", err)
			return
		}
		err = mgorm.Select(tx, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestCompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *mgorm.SelectStmt
		ActualStmt    *mgorm.SelectStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Select(nil, "column1").From("table").(*mgorm.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*mgorm.SelectStmt),
			mgorm.ErrInvalidMockExpectation,
		},
		{
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*mgorm.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 100).(*mgorm.SelectStmt),
			mgorm.ErrInvalidMockExpectation,
		},
	}

	for _, testCase := range testCases {
		err := mgorm.CompareStmts(testCase.ExpectedStmt, testCase.ActualStmt)

		// Validate if the expected error was occurred.
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}
