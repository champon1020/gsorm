package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestMockDB_Return(t *testing.T) {
	testCases := []struct {
		MockDBExpected   []*mgorm.Stmt
		WillReturn       interface{}
		ResultWillReturn map[int]interface{}
	}{
		{
			[]*mgorm.Stmt{},
			10001,
			map[int]interface{}{},
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			"some_name",
			map[int]interface{}{0: "some_name"},
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
				mgorm.Select(nil, "emp_no").From("employees").(*mgorm.Stmt),
			},
			10001,
			map[int]interface{}{1: 10001},
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockDB)
		mock.ExportedSetExpected(testCase.MockDBExpected)
		mock.ExportedSetWillReturn(make(map[int]interface{}))

		mock.Return(testCase.WillReturn)
		if diff := cmp.Diff(testCase.ResultWillReturn, mock.ExportedGetWillReturn()); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestMockDB_CompareTo(t *testing.T) {
	testCases := []struct {
		MockDBExpected   []*mgorm.Stmt
		MockDBWillReturn map[int]interface{}
		Executed         *mgorm.Stmt
		ResultReturned   interface{}
	}{
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			map[int]interface{}{},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			nil,
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			map[int]interface{}{0: "some_name"},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
				mgorm.Select(nil, "emp_no").From("employees").(*mgorm.Stmt),
			},
			map[int]interface{}{0: "some_name"},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockDB)
		mock.ExportedSetExpected(testCase.MockDBExpected)
		mock.ExportedSetWillReturn(testCase.MockDBWillReturn)

		returned, err := mgorm.MockDBCompareTo(mock, testCase.Executed)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.ResultReturned, returned); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
