package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStmt_ProcessQuerySQL(t *testing.T) {
	testCases := []struct {
		Stmt   *Stmt
		Result string
	}{
		{
			&Stmt{
				Cmd:  &Select{Columns: []Column{{Name: "column"}}},
				From: &From{Tables: []Table{{Name: "table"}}},
			},
			"SELECT column FROM table",
		},
		{
			&Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				From:      &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs = ?", Values: []interface{}{10}},
			},
			"SELECT column FROM table WHERE lhs = 10",
		},
		{
			&Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				From:      &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&Stmt{
				Cmd:       &Select{Columns: []Column{{Name: "column"}}},
				From:      &From{Tables: []Table{{Name: "table"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&Or{Expr: "lhs2 = ? AND lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"SELECT column FROM table WHERE lhs1 = 10 OR (lhs2 = 20 AND lhs3 = 30)",
		},
	}

	for _, testCase := range testCases {
		sql, _ := testCase.Stmt.processQuerySQL()
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
				Cmd: &Insert{
					Table:   Table{Name: "table"},
					Columns: []Column{{Name: "column1"}, {Name: "column2"}},
				},
				Values: &Values{Columns: []interface{}{10, 20}},
			},
			"INSERT INTO table (column1, column2) VALUES (10, 20)",
		},
		{
			&Stmt{
				Cmd: &Update{Table: Table{Name: "table"}},
				Set: &Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2",
		},
		{
			&Stmt{
				Cmd:       &Update{Table: Table{Name: "table"}},
				Set:       &Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
				WhereExpr: &Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				AndOr: []Expr{
					&And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
				},
			},
			"UPDATE table SET lhs1 = rhs1, lhs2 = rhs2 WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&Stmt{
				Cmd:  &Delete{},
				From: &From{Tables: []Table{{Name: "table"}}},
			},
			"DELETE FROM table",
		},
	}

	for _, testCase := range testCases {
		sql, _ := testCase.Stmt.processExecSQL()
		assert.Equal(t, testCase.Result, string(sql))
	}
}
