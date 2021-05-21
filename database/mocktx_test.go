package database_test

import (
	"database/sql"
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/database"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
	"gotest.tools/v3/assert"
)

func TestMockTx_GetDriver(t *testing.T) {
	mock := database.NewMockDB("mysql")
	mock.ExpectBegin()
	mocktx, err := mock.Begin()
	if err != nil {
		t.Errorf("error was occurred: %v", err)
	}

	assert.Equal(t, database.MysqlDriver, mocktx.GetDriver())
}

func TestMockTx_DummyFunctions(t *testing.T) {
	mock := database.NewMockDB("")
	mock.ExpectBegin()

	mocktx, err := mock.Begin()
	if err != nil {
		t.Errorf("error was occurred: %v", err)
	}

	assert.Equal(t, nil, mocktx.Ping())

	r, e := mocktx.Exec("")
	assert.Equal(t, nil, r)
	assert.Equal(t, nil, e)

	r2, e2 := mocktx.Query("")
	var rexpected *sql.Rows
	assert.Equal(t, rexpected, r2)
	assert.Equal(t, nil, e2)
}

func TestMock_TransactionExpectation(t *testing.T) {
	expectedReturn1 := []int{10, 20, 30}
	expectedReturn2 := []string{"hello", "world", "!"}

	// Test phase.
	mock := database.NewMockDB("")
	mocktx1 := mock.ExpectBegin()
	mocktx2 := mock.ExpectBegin()

	mocktx1.Expect(gsorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mocktx1.ExpectWithReturn(gsorm.Select(nil, "column1").From("table1"), expectedReturn1)
	mocktx1.ExpectCommit()
	mocktx2.Expect(gsorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))
	mocktx2.ExpectWithReturn(gsorm.Select(nil, "column2").From("table2"), expectedReturn2)
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

	if err := gsorm.Insert(tx1, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	model1 := new([]int)
	if err := gsorm.Select(tx1, "column1").From("table1").Query(model1); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := tx1.Commit(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	if err := gsorm.Insert(tx2, "table2", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}

	model2 := new([]string)
	if err := gsorm.Select(tx2, "column2").From("table2").Query(model2); err != nil {
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

func TestMockTx_Commit_Fail(t *testing.T) {
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %+v", err)
			return
		}
		err = gsorm.Select(tx, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		mocktx := mock.ExpectBegin()
		mocktx.ExpectCommit()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %+v", err)
			return
		}
		err = gsorm.Select(tx, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		mocktx := mock.ExpectBegin()
		mocktx.ExpectRollback()

		// Actual process.
		model := new([]int)
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occurred: %+v", err)
			return
		}
		err = gsorm.Select(tx, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockTx_CompareWith_Fail(t *testing.T) {
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		mock.ExpectBegin()
		mocktx, err := mock.Begin()
		if err != nil {
			t.Errorf("error was occurred: %v", err)
		}
		mocktx.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err = gsorm.Insert(mocktx, "table2", "column2").Values(10).Exec()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		mock.ExpectBegin()
		mocktx, err := mock.Begin()
		if err != nil {
			t.Errorf("error was occurred: %v", err)
		}
		mocktx.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err = gsorm.Insert(mocktx, "table2", "column2").Values(10).Values(100).Exec()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}
