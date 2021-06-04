package statement_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/google/go-cmp/cmp"
)

func TestStatement_QueryWithMock(t *testing.T) {
	type Employee struct {
		EmpNo     int
		FirstName string
	}
	model := []Employee{}

	mock := gsorm.OpenMock()
	expectedReturn := []Employee{
		{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"},
	}
	mock.ExpectWithReturn(gsorm.Select(mock, "emp_no", "first_name").From("employees"), expectedReturn)

	err := gsorm.Select(mock, "emp_no", "first_name").From("employees").Query(&model)
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if diff := cmp.Diff(expectedReturn, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestStatement_ExecWithMock(t *testing.T) {
	mock := gsorm.OpenMock()
	mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))

	err := gsorm.Insert(mock, "employees", "emp_no", "first_name").Values(1001, "Taro").Exec()
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}
}
