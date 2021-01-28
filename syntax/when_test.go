package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhen_String(t *testing.T) {
	testCases := []struct {
		When   *syntax.When
		Result string
	}{
		{
			&syntax.When{Expr: "lhs = rhs"},
			`WHEN("lhs = rhs")`,
		},
		{
			&syntax.When{Expr: "lhs = ?", Values: []interface{}{10}},
			`WHEN("lhs = ?", 10)`,
		},
		{
			&syntax.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
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
		When   *syntax.When
		Result *syntax.StmtSet
	}{
		{
			&syntax.When{Expr: "lhs = rhs"},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = rhs"},
		},
		{
			&syntax.When{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = 10"},
		},
		{
			&syntax.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
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
		Result *syntax.When
	}{
		{
			"lhs = rhs",
			nil,
			&syntax.When{Expr: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{10},
			&syntax.When{Expr: "lhs = ?", Values: []interface{}{10}},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&syntax.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewWhen(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
