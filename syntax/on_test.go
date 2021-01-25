package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOn_Name(t *testing.T) {
	o := new(syntax.On)
	assert.Equal(t, "ON", syntax.OnName(o))
}

func TestOn_Build(t *testing.T) {
	testCases := []struct {
		On     *syntax.On
		Result *syntax.StmtSet
	}{
		{
			&syntax.On{Expr: "lhs = rhs"},
			&syntax.StmtSet{Clause: "ON", Value: "lhs = rhs"},
		},
		{
			&syntax.On{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "ON", Value: "lhs = 10"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.On.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOn(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.On
	}{
		{
			"lhs = rhs",
			nil,
			&syntax.On{Expr: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{10},
			&syntax.On{Expr: "lhs = ?", Values: []interface{}{10}},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&syntax.On{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewOn(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
