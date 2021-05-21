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

func TestNewMockDB(t *testing.T) {
	{
		var expected database.ExportedMockDB
		expected.ExportedSetDriver(database.MysqlDriver)
		mock := database.NewMockDB("mysql")
		assert.Equal(t, expected.GetDriver(), mock.GetDriver())
	}
	{
		var expected database.ExportedMockDB
		expected.ExportedSetDriver(database.PsqlDriver)
		mock := database.NewMockDB("psql")
		assert.Equal(t, expected.GetDriver(), mock.GetDriver())
	}
}

func TestMock_Expectation(t *testing.T) {
	expectedReturn := []int{10, 20, 30}

	// Test phase.
	mock := database.NewMockDB("")
	mock.Expect(gsorm.Insert(nil, "table", "column1", "column2").Values(10, "str"))
	mock.ExpectWithReturn(gsorm.Select(nil, "column1").From("table"), expectedReturn)

	// Actual process.
	if err := gsorm.Insert(mock, "table", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	model := new([]int)
	if err := gsorm.Select(mock, "column1").From("table").Query(model); err != nil {
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

func TestMockDB_DummyFunctions(t *testing.T) {
	mock := database.NewMockDB("")
	assert.Equal(t, nil, mock.Ping())
	assert.Equal(t, nil, mock.SetConnMaxLifetime(0))
	assert.Equal(t, nil, mock.SetMaxIdleConns(0))
	assert.Equal(t, nil, mock.SetMaxOpenConns(0))
	assert.Equal(t, nil, mock.Close())

	r, e := mock.Exec("")
	assert.Equal(t, nil, r)
	assert.Equal(t, nil, e)

	r2, e2 := mock.Query("")
	var rexpected *sql.Rows
	assert.Equal(t, rexpected, r2)
	assert.Equal(t, nil, e2)
}

func TestMockDB_Begin_Fail(t *testing.T) {
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := new(database.ExportedMockDB)

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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := new(database.ExportedMockDB)
		mock.Expect(gsorm.Select(nil, "column1").From("table"))
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
	expectedErr := database.ErrInvalidMockExpectation

	// Test phase.
	mock := database.NewMockDB("")
	mock.Expect(gsorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mock.Expect(gsorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	if err := gsorm.Insert(mock, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
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
	expectedErr := database.ErrInvalidMockExpectation

	// Test phase.
	mock := database.NewMockDB("")
	mocktx := mock.ExpectBegin()
	mocktx.Expect(gsorm.Insert(nil, "table1", "column1", "column2").Values(10, "str"))
	mocktx.Expect(gsorm.Insert(nil, "table2", "column1", "column2").Values(10, "str"))

	// Actual process.
	tx, err := mock.Begin()
	if err != nil {
		t.Errorf("Error was occured: %+v", err)
		return
	}
	if err = gsorm.Insert(tx, "table1", "column1", "column2").Values(10, "str").Exec(); err != nil {
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
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")

		// Actual process.
		model := new([]int)
		err := gsorm.Select(mock, "column1").From("table").Query(model)

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
		_ = mock.ExpectBegin()

		// Actual process.
		model := new([]int)
		err := gsorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestMockDB_CompareWith_Fail(t *testing.T) {
	{
		expectedErr := database.ErrInvalidMockExpectation

		// Test phase.
		mock := database.NewMockDB("")
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table2", "column2").Values(10).Exec()

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
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table2", "column2").Values(10).Values(100).Exec()

		// Validate if the expected error was occurred.
		if !failure.Is(err, expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", expectedErr)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}
