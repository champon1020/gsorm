package mgorm_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/champon1020/mgorm/syntax/expr"
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
			&cmd.Select{},
			[]syntax.Expr{
				&expr.From{},
				&expr.Where{},
			},
			&cmd.Select{},
			[]syntax.Expr{
				&expr.From{},
				&expr.Where{},
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
			&cmd.Select{},
			[]syntax.Expr{
				&expr.From{},
				&expr.Where{},
			},
			&cmd.Select{},
			[]syntax.Expr{
				&expr.From{},
				&expr.Where{},
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

func TestMockDB_Result(t *testing.T) {
	testCases := []struct {
		Expected []*mgorm.Stmt
		Actual   []*mgorm.Stmt
	}{
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).(*mgorm.Stmt),
			},
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).ExpectQuery(nil),
			},
		},
	}

	for _, testCase := range testCases {
		mockdb := new(mgorm.MockDB)
		mockdb.ExportedSetExpected(testCase.Expected)
		mockdb.ExportedSetActual(testCase.Actual)
		if err := mockdb.Result(); err != nil {
			t.Errorf("Error was occurred:\n  %v", err)
		}
	}
}

func TestMockDB_Result_Fail(t *testing.T) {
	testCases := []struct {
		Expected []*mgorm.Stmt
		Actual   []*mgorm.Stmt
		Error    error
	}{
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).(*mgorm.Stmt),
			},
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).ExpectQuery(nil),
				mgorm.Select(nil, "*").From("table").Where("lhs1 = ? AND lhs2 = ?", 10, "str").ExpectQuery(nil),
			},
			errors.New(`SELECT("*").FROM("table").WHERE("lhs1 = ? AND lhs2 = ?", 10, "str") was executed, but not expected`),
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).(*mgorm.Stmt),
				mgorm.Select(nil, "*").From("table").Where("lhs1 = ? AND lhs2 = ?", 10, "str").(*mgorm.Stmt),
			},
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).ExpectQuery(nil),
			},
			errors.New(`no query was executed, but SELECT("*").FROM("table").WHERE("lhs1 = ? AND lhs2 = ?", 10, "str") is expected`),
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs1 = ? AND lhs2 = ?", 10, "str").(*mgorm.Stmt),
			},
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).ExpectQuery(nil),
			},
			errors.New(`SELECT("*").FROM("table").WHERE("lhs = ?", 10) was executed, but SELECT("*").FROM("table").WHERE("lhs1 = ? AND lhs2 = ?", 10, "str") is expected`),
		},
		{
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs = ?", 10).ExpectQuery(nil),
			},
			[]*mgorm.Stmt{
				mgorm.Select(nil, "*").From("table").Where("lhs1 = ? AND lhs2 = ?", 10, "str").(*mgorm.Stmt),
			},
			errors.New(`SELECT("*").FROM("table").WHERE("lhs1 = ? AND lhs2 = ?", 10, "str") was executed, but SELECT("*").FROM("table").WHERE("lhs = ?", 10) is expected`),
		},
	}

	for _, testCase := range testCases {
		mockdb := new(mgorm.MockDB)
		mockdb.ExportedSetExpected(testCase.Expected)
		mockdb.ExportedSetActual(testCase.Actual)
		err := mockdb.Result()
		if err == nil {
			t.Errorf("Error was not occurred\n")
		}
		if err.Error() != testCase.Error.Error() {
			t.Errorf("\nGot : %v\nWant: %v\n", err, testCase.Error)
		}
	}
}
