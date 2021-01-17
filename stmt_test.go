package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStmt_ProcessQuerySQL(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result string
	}{
		{
			&mgorm.Stmt{
				Cmd:      &Select{Columns: []Column{{Name: "column"}}},
				FromExpr: &From{Tables: []Table{{Name: "table"}}},
			},
			"SELECT column FROM table",
		},
		{
			&mgorm.Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				FromExpr:  &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs = ?", Values: []interface{}{10}},
			},
			"SELECT column FROM table WHERE lhs = 10",
		},
		{
			&mgorm.Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				FromExpr:  &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&mgorm.Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				FromExpr:  &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&Or{Expr: "lhs2 = ? AND lhs3 = ?", Values: []interface{}{20, 30}},
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
		Stmt   *mgorm.Stmt
		Result string
	}{
		{
			&mgorm.Stmt{
				Cmd: &Insert{
					Table:   Table{Name: "table"},
					Columns: []Column{{Name: "column1"}, {Name: "column2"}},
				},
				ValuesExpr: &Values{Columns: []interface{}{10, 20}},
			},
			"INSERT INTO table (column1, column2) VALUES (10, 20)",
		},
		{
			&mgorm.Stmt{
				Cmd:     &Update{Table: Table{Name: "table"}},
				SetExpr: &Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2",
		},
		{
			&mgorm.Stmt{
				Cmd:       &Update{Table: Table{Name: "table"}},
				SetExpr:   &Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2 WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&mgorm.Stmt{
				Cmd:      &Delete{},
				FromExpr: &From{Tables: []Table{{Name: "table"}}},
			},
			"DELETE FROM table",
		},
	}

	for _, testCase := range testCases {
		sql, _ := StmtProcessExecSQL(testCase.Stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}

func TestStmt_Where(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *mgorm.Stmt
		Result *mgorm.Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&mgorm.Stmt{},
			&mgorm.Stmt{WhereExpr: &syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}}},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.Where(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestStmt_And(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *mgorm.Stmt
		Result *mgorm.Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&mgorm.Stmt{},
			&mgorm.Stmt{AndOr: []syntax.Expr{&syntax.And{Expr: "lhs = ?", Values: []interface{}{10}}}},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.And(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestStmt_Or(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Stmt   *mgorm.Stmt
		Result *mgorm.Stmt
	}{
		{
			"lhs = ?",
			[]interface{}{10},
			&mgorm.Stmt{},
			&mgorm.Stmt{AndOr: []syntax.Expr{&syntax.Or{Expr: "lhs = ?", Values: []interface{}{10}}}},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.Or(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
