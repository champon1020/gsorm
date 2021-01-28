package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhere_String(t *testing.T) {
	testCases := []struct {
		Where  *expr.Where
		Result string
	}{
		{
			&expr.Where{Expr: "lhs = rhs"},
			`WHERE("lhs = rhs")`,
		},
		{
			&expr.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			`WHERE("lhs = ?", 10)`,
		},
		{
			&expr.Where{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
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
		Where  *expr.Where
		Result *syntax.StmtSet
	}{
		{
			&expr.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHERE", Value: "lhs = 10"},
		},
		{
			&expr.Where{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
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
		Result *expr.Where
	}{
		{"lhs = ?", []interface{}{"rhs"}, &expr.Where{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := expr.NewWhere(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
