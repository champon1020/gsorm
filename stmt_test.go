package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/syntax"
	"github.com/stretchr/testify/assert"
)

func TestStmt_ProcessQuerySQL(t *testing.T) {
	testCases := []struct {
		Cmd    *syntax.Select
		Called []syntax.Expr
		Result string
	}{
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
			[]syntax.Expr{
				&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"SELECT column FROM table",
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
			[]syntax.Expr{
				&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				&syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			},
			"SELECT column FROM table WHERE lhs = 10",
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
			[]syntax.Expr{
				&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				&syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
			},
			"SELECT column FROM table WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)",
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
			[]syntax.Expr{
				&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
				&syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				&syntax.Or{Expr: "lhs2 = ? AND lhs3 = ?", Values: []interface{}{20, 30}},
			},
			"SELECT column FROM table WHERE lhs1 = 10 OR (lhs2 = 20 AND lhs3 = 30)",
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		stmt.ExportedSetCalled(testCase.Called)
		sql, _ := mgorm.StmtProcessQuerySQL(stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}

func TestStmt_ProcessCaseSQL(t *testing.T) {
	testCases := []struct {
		Called []syntax.Expr
		Result string
	}{
		{
			[]syntax.Expr{
				&syntax.When{Expr: "lhs > rhs"},
				&syntax.Then{"value"},
			},
			`CASE WHEN lhs > rhs THEN "value" END`,
		},
		{
			[]syntax.Expr{
				&syntax.When{Expr: "lhs1 > rhs1"},
				&syntax.Then{"value1"},
				&syntax.When{Expr: "lhs2 < rhs2"},
				&syntax.Then{"value2"},
			},
			`CASE WHEN lhs1 > rhs1 THEN "value1" WHEN lhs2 < rhs2 THEN "value2" END`,
		},
		{
			[]syntax.Expr{
				&syntax.When{Expr: "lhs1 > rhs1"},
				&syntax.Then{"value1"},
				&syntax.When{Expr: "lhs2 < rhs2"},
				&syntax.Then{"value2"},
				&syntax.Else{"value3"},
			},
			`CASE WHEN lhs1 > rhs1 THEN "value1" WHEN lhs2 < rhs2 THEN "value2" ELSE "value3" END`,
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCalled(testCase.Called)
		sql, _ := mgorm.StmtProcessCaseSQL(stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}

func TestStmt_ProcessExecSQL(t *testing.T) {
	testCases := []struct {
		Cmd    syntax.Cmd
		Called []syntax.Expr
		Result string
	}{
		{
			&syntax.Insert{
				Table:   syntax.Table{Name: "table"},
				Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}},
			},
			[]syntax.Expr{
				&syntax.Values{Columns: []interface{}{10, 20}},
			},
			"INSERT INTO table (column1, column2) VALUES (10, 20)",
		},
		{
			&syntax.Update{Table: syntax.Table{Name: "table"}},
			[]syntax.Expr{
				&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			},
			`UPDATE table SET lhs1 = "rhs1", lhs2 = "rhs2"`,
		},
		{
			&syntax.Update{Table: syntax.Table{Name: "table"}},
			[]syntax.Expr{
				&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
				&syntax.Where{Expr: "lhs1 = ?", Values: []interface{}{10}},
				&syntax.And{Expr: "lhs2 = ? OR lhs3 = ?", Values: []interface{}{20, 30}},
			},
			`UPDATE table SET lhs1 = "rhs1", lhs2 = "rhs2" WHERE lhs1 = 10 AND (lhs2 = 20 OR lhs3 = 30)`,
		},
		{
			&syntax.Delete{},
			[]syntax.Expr{
				&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			},
			"DELETE FROM table",
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		stmt.ExportedSetCalled(testCase.Called)
		sql, _ := mgorm.StmtProcessExecSQL(stmt)
		assert.Equal(t, testCase.Result, string(sql))
	}
}
