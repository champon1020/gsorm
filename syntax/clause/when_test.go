package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhen_String(t *testing.T) {
	testCases := []struct {
		When   *clause.When
		Result string
	}{
		{
			&clause.When{Expr: "lhs = rhs"},
			`WHEN("lhs = rhs")`,
		},
		{
			&clause.When{Expr: "lhs = ?", Values: []interface{}{10}},
			`WHEN("lhs = ?", 10)`,
		},
		{
			&clause.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
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
		When   *clause.When
		Result *syntax.StmtSet
	}{
		{
			&clause.When{Expr: "lhs = rhs"},
			&syntax.StmtSet{Keyword: "WHEN", Value: "lhs = rhs"},
		},
		{
			&clause.When{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Keyword: "WHEN", Value: "lhs = 10"},
		},
		{
			&clause.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Keyword: "WHEN", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.When.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewWhen(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *clause.When
	}{
		{
			"lhs = rhs",
			nil,
			&clause.When{Expr: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{10},
			&clause.When{Expr: "lhs = ?", Values: []interface{}{10}},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&clause.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewWhen(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
