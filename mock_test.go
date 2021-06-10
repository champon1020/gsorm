package gsorm_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestMock_Expectation(t *testing.T) {
	expectedReturn := []int{10, 20, 30}

	// Test phase.
	mock := gsorm.OpenMock()
	mock.Expect(gsorm.Insert(nil, "table", "column1", "column2").Values(10, "str"))
	mock.ExpectWithReturn(gsorm.Select(nil, "column1").From("table"), expectedReturn)

	// Actual process.
	if err := gsorm.Insert(mock, "table", "column1", "column2").Values(10, "str").Exec(); err != nil {
		t.Errorf("Error was occurred: %+v", err)
		return
	}
	model := &[]int{}
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
	mock := gsorm.OpenMock()
	assert.Equal(t, nil, mock.Ping())
	assert.Equal(t, nil, mock.SetConnMaxLifetime(0))
	assert.Equal(t, nil, mock.SetMaxIdleConns(0))
	assert.Equal(t, nil, mock.SetMaxOpenConns(0))
	assert.Equal(t, nil, mock.Close())

	r, e := mock.Exec("")
	assert.Equal(t, nil, r)
	assert.Equal(t, nil, e)

	r2, e2 := mock.Query("")
	var rexpected gsorm.ExportedIRows
	assert.Equal(t, rexpected, r2)
	assert.Equal(t, nil, e2)
}

func TestMockDB_Begin_Fail(t *testing.T) {
	{
		expectedErr := "gsorm.mockDB.Begin is not expected"

		// Test phase.
		mock := &gsorm.ExportedMockDB{}

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "gsorm.mockDB.Begin is not expected"

		// Test phase.
		mock := &gsorm.ExportedMockDB{}
		mock.Expect(gsorm.Select(nil, "column1").From("table"))
		_ = mock.ExpectBegin()

		// Actual process.
		_, err := mock.Begin()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
}

func TestMockDB_Complete_Fail(t *testing.T) {
	expectedErr := `Insert("table2", "column1", "column2").Values(10, "str") is expected but not executed`

	// Test phase.
	mock := gsorm.OpenMock()
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
	assert.EqualError(t, err, expectedErr)
}

func TestMockDB_Complete_Transaction_Fail(t *testing.T) {
	expectedErr := `Insert("table2", "column1", "column2").Values(10, "str") is expected but not executed`

	// Test phase.
	mock := gsorm.OpenMock()
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
	assert.EqualError(t, err, expectedErr)
}

func TestMockDB_CompareWith(t *testing.T) {
	{
		expectedErr := `Select("column1").From("table") is not expected but executed`

		// Test phase.
		mock := gsorm.OpenMock()

		// Actual process.
		model := &[]int{}
		err := gsorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\n" +
			"expected: gsorm.MockDB.Begin\n" +
			"actual:   Select(\"column1\").From(\"table\")\n"

		// Test phase.
		mock := gsorm.OpenMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := &[]int{}
		err := gsorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\n" +
			"expected: Insert(\"table1\", \"column1\").Values(10)\n" +
			"actual:   Insert(\"table2\", \"column2\").Values(10)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table2", "column2").Values(10).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\n" +
			"expected: Insert(\"table1\", \"column1\").Values(10)\n" +
			"actual:   Insert(\"table1\", \"column1\").Values(10).Values(100)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table1", "column1").Values(10).Values(100).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
}

func TestMockTx_DummyFunctions(t *testing.T) {
	mock := gsorm.OpenMock()
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
	var rexpected gsorm.ExportedIRows
	assert.Equal(t, rexpected, r2)
	assert.Equal(t, nil, e2)
}

func TestMock_TransactionExpectation(t *testing.T) {
	expectedReturn1 := []int{10, 20, 30}
	expectedReturn2 := []string{"hello", "world", "!"}

	// Test phase.
	mock := gsorm.OpenMock()
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
		expectedErr := `gsorm.mockTx.Commit is not expected`

		// Test phase.
		mock := gsorm.OpenMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Commit()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := `gsorm.mockTx.Commit is not expected`

		// Test phase.
		mock := gsorm.OpenMock()
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
		assert.EqualError(t, err, expectedErr)
	}
}

func TestMockTx_Rollback_Fail(t *testing.T) {
	{
		expectedErr := `gsorm.mockTx.Rollback is not expected`

		// Test phase.
		mock := gsorm.OpenMock()
		_ = mock.ExpectBegin()

		// Actual process.
		tx, err := mock.Begin()
		if err != nil {
			t.Errorf("Error was occured: %+v", err)
			return
		}
		err = tx.Rollback()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := `gsorm.mockTx.Rollback is not expected`

		// Test phase.
		mock := gsorm.OpenMock()
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
		assert.EqualError(t, err, expectedErr)
	}
}

func TestMockTx_CompareWith(t *testing.T) {
	{
		expectedErr := `Select("column1").From("table") is not expected but executed`

		// Test phase.
		mock := gsorm.OpenMock()
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
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\n" +
			"expected: gsorm.MockTx.Commit\n" +
			"actual:   Select(\"column1\").From(\"table\")\n"

		// Test phase.
		mock := gsorm.OpenMock()
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
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\n" +
			"expected: gsorm.MockTx.Rollback\n" +
			"actual:   Select(\"column1\").From(\"table\")\n"

		// Test phase.
		mock := gsorm.OpenMock()
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
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\nexpected: Insert(\"table1\", \"column1\").Values(10)\nactual:   Insert(\"table2\", \"column2\").Values(10)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mocktx := mock.ExpectBegin()
		mocktx.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mocktx, "table2", "column2").Values(10).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\nexpected: Insert(\"table1\", \"column1\").Values(10)\nactual:   Insert(\"table1\", \"column1\").Values(10).Values(100)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mocktx := mock.ExpectBegin()
		mocktx.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mocktx, "table1", "column1").Values(10).Values(100).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
}
