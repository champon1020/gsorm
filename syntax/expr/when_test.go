package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhen_String(t *testing.T) {
	testCases := []struct {
		When   *expr.When
		Result string
	}{
		{
			&expr.When{Expr: "lhs = rhs"},
			`WHEN("lhs = rhs")`,
		},
		{
			&expr.When{Expr: "lhs = ?", Values: []interface{}{10}},
			`WHEN("lhs = ?", 10)`,
		},
		{
			&expr.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`WHEN("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.When.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestWhen_Build(t *testing.T) {
	testCases := []struct {
		When   *expr.When
		Result *syntax.StmtSet
	}{
		{
			&expr.When{Expr: "lhs = rhs"},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = rhs"},
		},
		{
			&expr.When{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = 10"},
		},
		{
			&expr.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Clause: "WHEN", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.When.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewWhen(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *expr.When
	}{
		{
			"lhs = rhs",
			nil,
			&expr.When{Expr: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{10},
			&expr.When{Expr: "lhs = ?", Values: []interface{}{10}},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&expr.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewWhen(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
