package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/syntax"
	"github.com/stretchr/testify/assert"
)

func TestStmt_ProcessQuerySQL(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result string
	}{
		{
			&mgorm.Stmt{
				Cmd:      &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				FromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"SELECT column FROM table",
		},
		{
			&mgorm.Stmt{
				Cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				FromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				WhereExpr: &syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			},
			"SELECT column FROM table WHERE lhs = 10",
		},
		{
			&mgorm.Stmt{
				Cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				FromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				WhereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []syntax.Expr{
					&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&mgorm.Stmt{
				Cmd:       &syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
				FromExpr:  &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				WhereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []syntax.Expr{
					&syntax.Or{Expr: "lhs2 = ? AND lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 OR (lhs2 = 20 AND lhs3 = 30)",
		},
	}

	for _, testCase := range testCases {
		sql, _ := mgorm.StmtProcessQuerySQL(testCase.Stmt)
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
				Cmd: &syntax.Insert{
					Table:   syntax.Table{Name: "table"},
					Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}},
				},
				ValuesExpr: &syntax.Values{Columns: []interface{}{10, 20}},
			},
			"INSERT INTO table (column1, column2) VALUES (10, 20)",
		},
		{
			&mgorm.Stmt{
				Cmd:     &syntax.Update{Table: syntax.Table{Name: "table"}},
				SetExpr: &syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2",
		},
		{
			&mgorm.Stmt{
				Cmd:       &syntax.Update{Table: syntax.Table{Name: "table"}},
				SetExpr:   &syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
				WhereExpr: &syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []syntax.Expr{
					&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2 WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&mgorm.Stmt{
				Cmd:      &syntax.Delete{},
				FromExpr: &syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"DELETE FROM table",
		},
	}

	for _, testCase := range testCases {
		sql, _ := mgorm.StmtProcessExecSQL(testCase.Stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}
