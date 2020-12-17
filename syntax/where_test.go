package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestWhere_Build(t *testing.T) {
	testCases := []struct {
		Where  *syntax.Where
		Result *syntax.StmtSet
	}{
		{
			&syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHERE", Value: "lhs = 10"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Where.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
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
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestAnd_Build(t *testing.T) {
	testCases := []struct {
		And    *syntax.And
		Result *syntax.StmtSet
	}{
		{
			&syntax.And{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "AND", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.And.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewAdd(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.And
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.And{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewAnd(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestOr_Build(t *testing.T) {
	testCases := []struct {
		Or     *syntax.Or
		Result *syntax.StmtSet
	}{
		{
			&syntax.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "OR", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Or.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOr(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.Or
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.Or{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewOr(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
