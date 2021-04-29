package database_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/database"
	"github.com/google/go-cmp/cmp"
)

func TestMock_TransactionExpectation(t *testing.T) {
	expectedReturn1 := []int{10, 20, 30}
	expectedReturn2 := []string{"hello", "world", "!"}

	// Test phase.
	mock := database.NewMockDB()
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
