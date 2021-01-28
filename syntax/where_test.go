package syntax_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhere_String(t *testing.T) {
	testCases := []struct {
		Where  *syntax.Where
		Result string
	}{
		{
			&syntax.Where{Expr: "lhs = rhs"},
			`WHERE("lhs = rhs")`,
		},
		{
			&syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			`WHERE("lhs = ?", 10)`,
		},
		{
			&syntax.Where{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`WHERE("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Where.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestWhere_Build(t *testing.T) {
	testCases := []struct {
		Where  *syntax.Where
		Result *syntax.StmtSet
	}{
		{
			&syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHERE", Value: "lhs = 10"},
		},
		{
			&syntax.Where{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Clause: "WHERE", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Where.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewWhere(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.Where
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.Where{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewWhere(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestBuildStmtSet(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.StmtSet
	}{
		{
			`lhs = "rhs"`,
			[]interface{}{},
			&syntax.StmtSet{Value: `lhs = "rhs"`},
		},
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: `lhs = "rhs"`},
		},
		{
			"lhs = ?",
			[]interface{}{100},
			&syntax.StmtSet{Value: "lhs = 100"},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{"rhs", 100},
			&syntax.StmtSet{Value: `lhs1 = "rhs" AND lhs2 = 100`},
		},
		{
			"IN lhs (?, ?, ?)",
			[]interface{}{"rhs", 100, true},
			&syntax.StmtSet{Value: `IN lhs ("rhs", 100, true)`},
		},
		{
			"lhs LIKE %%?%%",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: `lhs LIKE %"rhs"%`},
		},
		{
			"lhs BETWEEN ? AND ?",
			[]interface{}{10, 100},
			&syntax.StmtSet{Value: "lhs BETWEEN 10 AND 100"},
		},
	}

	for _, testCase := range testCases {
		res, _ := syntax.BuildStmtSet(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestBuildStmtSet_Fail(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Error  error
	}{
		{
			"lhs = ? AND rhs = ?",
			[]interface{}{10},
			internal.NewError(
				syntax.OpBuildStmtSet,
				internal.KindBasic,
				errors.New("Length of values is not valid"),
			),
		},
		{
			"lhs = ?",
			[]interface{}{[]string{}},
			internal.NewError(
				internal.OpToString,
				internal.KindType,
				errors.New("type is invalid"),
			),
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.BuildStmtSet(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(e, testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
	}
}
