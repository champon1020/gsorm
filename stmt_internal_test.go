package mgorm

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStmt_ProcessQuerySQL(t *testing.T) {
	testCases := []struct {
		Stmt   *Stmt
		Result string
	}{
		{
			&Stmt{
				cmd:      &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				fromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"SELECT column FROM table",
		},
		{
			&Stmt{
				cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				fromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				whereExpr: &syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			},
			"SELECT column FROM table WHERE lhs = 10",
		},
		{
			&Stmt{
				cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				fromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				whereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				andOr: []syntax.Expr{
					&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&Stmt{
				cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				fromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				whereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				andOr: []syntax.Expr{
					&syntax.Or{Expr: "lhs2 = ? AND lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 OR (lhs2 = 20 AND lhs3 = 30)",
		},
	}

	for _, testCase := range testCases {
		sql, _ := StmtProcessQuerySQL(testCase.Stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}

func TestStmt_PrcessExecSQL(t *testing.T) {
	testCases := []struct {
		Stmt   *Stmt
		Result string
	}{
		{
			&Stmt{
				cmd: &syntax.Insert{
					Table:   syntax.Table{Name: "table"},
					Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}},
				},
				valuesExpr: &syntax.Values{Columns: []interface{}{10, 20}},
			},
			"INSERT INTO table (column1, column2) VALUES (10, 20)",
		},
		{
			&Stmt{
				cmd:     &syntax.Update{Table: syntax.Table{Name: "table"}},
				setExpr: &syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			},
			`UPDATE table SET lhs1 = "rhs1", lhs2 = "rhs2"`,
		},
		{
			&Stmt{
				cmd:       &syntax.Update{Table: syntax.Table{Name: "table"}},
				setExpr:   &syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
				whereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				andOr: []syntax.Expr{
					&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			`UPDATE table SET lhs1 = "rhs1", lhs2 = "rhs2" WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)`,
		},
		{
			&Stmt{
				cmd:      &syntax.Delete{},
				fromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"DELETE FROM table",
		},
	}

	for _, testCase := range testCases {
		sql, _ := StmtProcessExecSQL(testCase.Stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}

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
				fromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table1"}, {Name: "table2", Alias: "t2"}}},
				called: []*opArgs{{
					op:   opFrom,
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
				whereExpr: &syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
				called: []*opArgs{{
					op:   opWhere,
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
				andOr: []syntax.Expr{&syntax.And{Expr: "lhs = ?", Values: []interface{}{10}}},
				called: []*opArgs{{
					op:   opAnd,
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
				andOr: []syntax.Expr{&syntax.Or{Expr: "lhs = ?", Values: []interface{}{10}}},
				called: []*opArgs{{
					op:   opOr,
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
