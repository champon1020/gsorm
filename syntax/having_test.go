package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestHaving_String(t *testing.T) {
	testCases := []struct {
		Having *syntax.Having
		Result string
	}{
		{
			&syntax.Having{Expr: "lhs = rhs"},
			`HAVING("lhs = rhs")`,
		},
		{
			&syntax.Having{Expr: "lhs = ?", Values: []interface{}{10}},
			`HAVING("lhs = ?", 10)`,
		},
		{
			&syntax.Having{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`HAVING("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Having.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestHaving_Build(t *testing.T) {
	testCases := []struct {
		Having *syntax.Having
		Result *syntax.StmtSet
	}{
		{
			&syntax.Having{Expr: "lhs = rhs"},
			&syntax.StmtSet{Clause: "HAVING", Value: "lhs = rhs"},
		},
		{
			&syntax.Having{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "HAVING", Value: "lhs = 10"},
		},
		{
			&syntax.Having{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Clause: "HAVING", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Having.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewHaving(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.Having
	}{
		{
			"lhs = rhs",
			nil,
			&syntax.Having{Expr: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{10},
			&syntax.Having{Expr: "lhs = ?", Values: []interface{}{10}},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&syntax.Having{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewHaving(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
