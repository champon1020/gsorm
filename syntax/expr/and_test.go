package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestAnd_String(t *testing.T) {
	testCases := []struct {
		And    *expr.And
		Result string
	}{
		{
			&expr.And{Expr: "lhs = rhs"},
			`AND("lhs = rhs")`,
		},
		{
			&expr.And{Expr: "lhs = ?", Values: []interface{}{10}},
			`AND("lhs = ?", 10)`,
		},
		{
			&expr.And{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`AND("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.And.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestAnd_Build(t *testing.T) {
	testCases := []struct {
		And    *expr.And
		Result *syntax.StmtSet
	}{
		{
			&expr.And{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "AND", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.And.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewAdd(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *expr.And
	}{
		{"lhs = ?", []interface{}{"rhs"}, &expr.And{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := expr.NewAnd(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
