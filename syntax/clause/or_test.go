package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOr_String(t *testing.T) {
	testCases := []struct {
		Or     *clause.Or
		Result string
	}{
		{
			&clause.Or{Expr: "lhs = rhs"},
			`OR("lhs = rhs")`,
		},
		{
			&clause.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			`OR("lhs = ?", 10)`,
		},
		{
			&clause.Or{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`OR("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Or.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOr_Build(t *testing.T) {
	testCases := []struct {
		Or     *clause.Or
		Result *syntax.StmtSet
	}{
		{
			&clause.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Keyword: "OR", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Or.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOr(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *clause.Or
	}{
		{"lhs = ?", []interface{}{"rhs"}, &clause.Or{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := clause.NewOr(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}