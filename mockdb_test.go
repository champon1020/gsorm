package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

func TestMockDB_ExpectReturn(t *testing.T) {
	testCases := []struct {
		Stmt       *mgorm.Stmt
		WillReturn interface{}
	}{
		{
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
		{
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			[]string{"some_name", "any_name"},
		},
		{
			mgorm.Select(nil, "emp_no", "first_name").From("employees").(*mgorm.Stmt),
			map[int]string{0: "some_name", 1: "any_name"},
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockDB)
		mock.ExportedPushExpected(testCase.Stmt, testCase.WillReturn)
		mock.Expect(testCase.Stmt).Return(testCase.WillReturn)

		expectation := mock.ExportedPopExpected()
		eq, ok := expectation.(*mgorm.ExpectedQuery)
		if !ok {
			t.Errorf("expectation is not ExportedQuery")
			continue
		}
		if diff := cmp.Diff(
			testCase.Stmt.ExportedGetCmd(),
			eq.ExportedGetStmt().ExportedGetCmd(),
		); diff != "" {
			internal.PrintTestDiff(t, diff)
			continue
		}
		if diff := cmp.Diff(
			testCase.Stmt.ExportedGetCalled(),
			eq.ExportedGetStmt().ExportedGetCalled(),
		); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestMockTx_ExpectReturn(t *testing.T) {
	testCases := []struct {
		Stmt       *mgorm.Stmt
		WillReturn interface{}
	}{
		{
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
		{
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			[]string{"some_name", "any_name"},
		},
		{
			mgorm.Select(nil, "emp_no", "first_name").From("employees").(*mgorm.Stmt),
			map[int]string{0: "some_name", 1: "any_name"},
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockTx)
		mock.ExportedPushExpected(testCase.Stmt, testCase.WillReturn)
		mock.Expect(testCase.Stmt).Return(testCase.WillReturn)

		expectation := mock.ExportedPopExpected()
		eq, ok := expectation.(*mgorm.ExpectedQuery)
		if !ok {
			t.Errorf("expectation is not ExportedQuery")
			continue
		}
		if diff := cmp.Diff(
			testCase.Stmt.ExportedGetCmd(),
			eq.ExportedGetStmt().ExportedGetCmd(),
		); diff != "" {
			internal.PrintTestDiff(t, diff)
			continue
		}
		if diff := cmp.Diff(
			testCase.Stmt.ExportedGetCalled(),
			eq.ExportedGetStmt().ExportedGetCalled(),
		); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestCompareTo_MockDB(t *testing.T) {
	testCases := []struct {
		ExpectedStmts       []*mgorm.Stmt
		ExpectedWillReturns []interface{}
		Executed            *mgorm.Stmt
		ResultReturned      interface{}
	}{
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				nil,
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			nil,
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				"some_name",
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
				mgorm.Select(nil, "emp_no").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				"some_name",
				"any_name",
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockDB)
		for i := 0; i < len(testCase.ExpectedStmts); i++ {
			mock.ExportedPushExpected(testCase.ExpectedStmts[i], testCase.ExpectedWillReturns[i])
		}

		returned, err := mgorm.CompareTo(mock, testCase.Executed)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.ResultReturned, returned); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestCompareTo_MockTx(t *testing.T) {
	testCases := []struct {
		ExpectedStmts       []*mgorm.Stmt
		ExpectedWillReturns []interface{}
		Executed            *mgorm.Stmt
		ResultReturned      interface{}
	}{
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				nil,
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			nil,
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				"some_name",
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
				mgorm.Select(nil, "emp_no").From("employees").(*mgorm.Stmt),
			},
			[]interface{}{
				"some_name",
				"any_name",
			},
			mgorm.Select(nil, "first_name").From("employees").(*mgorm.Stmt),
			"some_name",
		},
	}

	for _, testCase := range testCases {
		mock := new(mgorm.MockTx)
		for i := 0; i < len(testCase.ExpectedStmts); i++ {
			mock.ExportedPushExpected(testCase.ExpectedStmts[i], testCase.ExpectedWillReturns[i])
		}

		returned, err := mgorm.CompareTo(mock, testCase.Executed)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.ResultReturned, returned); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
