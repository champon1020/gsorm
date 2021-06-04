package gsorm_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/interfaces/domain"
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
	var rexpected domain.Rows
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
		//		expectedErr := "gsorm.mockDB.Begin is not expected"
		expectedErr := "statements comparison was failed:\nexpected: gsorm.MockDB.Begin\nactual:   Select(\"column1\").From(\"table\")\n"

		// Test phase.
		mock := gsorm.OpenMock()
		_ = mock.ExpectBegin()

		// Actual process.
		model := &[]int{}
		err := gsorm.Select(mock, "column1").From("table").Query(model)

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
}

func TestMockDB_CompareWith_Fail(t *testing.T) {
	{
		//		expectedErr := "gsorm.mockDB.Begin is not expected"
		expectedErr := "statements comparison was failed:\nexpected: Insert(\"table1\", \"column1\").Values(10)\nactual:   Insert(\"table2\", \"column2\").Values(10)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table2", "column2").Values(10).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
	{
		expectedErr := "statements comparison was failed:\nexpected: Insert(\"table1\", \"column1\").Values(10)\nactual:   Insert(\"table1\", \"column1\").Values(10).Values(100)\n"

		// Test phase.
		mock := gsorm.OpenMock()
		mock.Expect(gsorm.Insert(nil, "table1", "column1").Values(10))

		// Actual process.
		err := gsorm.Insert(mock, "table1", "column1").Values(10).Values(100).Exec()

		// Validate if the expected error was occurred.
		assert.EqualError(t, err, expectedErr)
	}
}
