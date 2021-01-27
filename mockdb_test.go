package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestMockDB_AddExecuted(t *testing.T) {
	testCases := []struct {
		MockDB       *mgorm.MockDB
		Cmd          syntax.Cmd
		Called       []syntax.Expr
		ResultCmd    syntax.Cmd
		ResultCalled []syntax.Expr
	}{
		{
			&mgorm.MockDB{},
			&syntax.Select{},
			[]syntax.Expr{
				&syntax.From{},
				&syntax.Where{},
			},
			&syntax.Select{},
			[]syntax.Expr{
				&syntax.From{},
				&syntax.Where{},
			},
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		stmt.ExportedSetCalled(testCase.Called)
		mgorm.MockDBAddExecuted(testCase.MockDB, stmt)

		actualCmd := testCase.MockDB.ExportedGetActual()[0].ExportedGetCmd()
		if diff := cmp.Diff(testCase.ResultCmd, actualCmd); diff != "" {
			internal.PrintTestDiff(t, diff)
		}

		actualCalled := testCase.MockDB.ExportedGetActual()[0].ExportedGetCalled()
		if diff := cmp.Diff(testCase.ResultCalled, actualCalled); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestMockDB_AddExpected(t *testing.T) {
	testCases := []struct {
		MockDB       *mgorm.MockDB
		Cmd          syntax.Cmd
		Called       []syntax.Expr
		ResultCmd    syntax.Cmd
		ResultCalled []syntax.Expr
	}{
		{
			&mgorm.MockDB{},
			&syntax.Select{},
			[]syntax.Expr{
				&syntax.From{},
				&syntax.Where{},
			},
			&syntax.Select{},
			[]syntax.Expr{
				&syntax.From{},
				&syntax.Where{},
			},
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		stmt.ExportedSetCalled(testCase.Called)
		testCase.MockDB.AddExpected(stmt)

		expectedCmd := testCase.MockDB.ExportedGetExpected()[0].ExportedGetCmd()
		if diff := cmp.Diff(testCase.ResultCmd, expectedCmd); diff != "" {
			internal.PrintTestDiff(t, diff)
		}

		expectedCalled := testCase.MockDB.ExportedGetExpected()[0].ExportedGetCalled()
		if diff := cmp.Diff(testCase.ResultCalled, expectedCalled); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
