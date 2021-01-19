package mgorm

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStmt_From(t *testing.T) {
	testCases := []struct {
		Tables []string
		Stmt   *Stmt
		Result *Stmt
	}{
		{
			[]string{"table1", "table2 AS t2"},
			&Stmt{},
			&Stmt{
				FromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table1"}, {Name: "table2", Alias: "t2"}}},
				called: []*opArgs{{
					op:   OpFrom,
					args: []interface{}{[]string{"table1", "table2 AS t2"}},
				}},
			},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.From(testCase.Tables...)
		if diff := cmp.Diff(res, testCase.Result, cmp.AllowUnexported(*res, opArgs{})); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
		assert.Equal(t, res.called, testCase.Result.called)
	}
}

func TestStmt_Where(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *Stmt
		Result *Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&Stmt{},
			&Stmt{
				WhereExpr: &syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
				called: []*opArgs{{
					op:   OpWhere,
					args: []interface{}{"lhs = ?", []interface{}{10}},
				}},
			},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.Where(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result, cmp.AllowUnexported(*res, opArgs{})); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
		assert.Equal(t, res.called, testCase.Result.called)
	}
}

func TestStmt_And(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *Stmt
		Result *Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&Stmt{},
			&Stmt{
				AndOr: []syntax.Expr{&syntax.And{Expr: "lhs = ?", Values: []interface{}{10}}},
				called: []*opArgs{{
					op:   OpAnd,
					args: []interface{}{"lhs = ?", []interface{}{10}},
				}},
			},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.And(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result, cmp.AllowUnexported(*res, opArgs{})); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
		assert.Equal(t, res.called, testCase.Result.called)
	}
}

func TestStmt_Or(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *Stmt
		Result *Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&Stmt{},
			&Stmt{
				AndOr: []syntax.Expr{&syntax.Or{Expr: "lhs = ?", Values: []interface{}{10}}},
				called: []*opArgs{{
					op:   OpOr,
					args: []interface{}{"lhs = ?", []interface{}{10}},
				}},
			},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.Or(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result, cmp.AllowUnexported(*res, opArgs{})); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
		assert.Equal(t, res.called, testCase.Result.called)
	}
}
