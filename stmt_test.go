package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/stretchr/testify/assert"
)

func TestStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result string
	}{
		{
			mgorm.Select(nil, "column1", "column2 AS c2").
				From("table1 AS t1").
				Join("table2 AS t2").
				On("t1.column1 = t2.column1").
				Where("lhs1 > ?", 10).
				And("lhs2 = ? OR lhs3 = ?", "str", true).
				OrderByDesc("t1.column1").
				Limit(5).
				Offset(6).
				Union(mgorm.Select(nil, "column1", "column2 AS c2").
					From("table3").
					Sub()).(*mgorm.Stmt),
			`SELECT column1, column2 AS c2 ` +
				`FROM table1 AS t1 ` +
				`INNER JOIN table2 AS t2 ` +
				`ON t1.column1 = t2.column1 ` +
				`WHERE lhs1 > 10 ` +
				`AND (lhs2 = "str" OR lhs3 = true) ` +
				`ORDER BY t1.column1 DESC ` +
				`LIMIT 5 ` +
				`OFFSET 6 ` +
				`UNION ` +
				`SELECT column1, column2 AS c2 ` +
				`FROM table3`,
		},
		{
			mgorm.Select(nil, "column1", "column2 AS c2").
				From("table1 AS t1").
				LeftJoin("table2 AS t2").
				On("t1.column1 = t2.column1").
				Where("lhs1 > ?", 10).
				Or("lhs2 = ? AND lhs3 = ?", "str", true).
				UnionAll(mgorm.Select(nil, "column1", "column2 AS c2").
					From("table3").
					Sub()).(*mgorm.Stmt),
			`SELECT column1, column2 AS c2 ` +
				`FROM table1 AS t1 ` +
				`LEFT JOIN table2 AS t2 ` +
				`ON t1.column1 = t2.column1 ` +
				`WHERE lhs1 > 10 ` +
				`OR (lhs2 = "str" AND lhs3 = true) ` +
				`UNION ALL ` +
				`SELECT column1, column2 AS c2 ` +
				`FROM table3`,
		},
		{
			mgorm.Select(nil, "COUNT(column1)").
				From("table1 AS t1").
				RightJoin("table2 AS t2").
				On("t1.column1 = t2.column1").
				GroupBy("column1").
				Having("column1 > ?", 10).(*mgorm.Stmt),
			`SELECT COUNT(column1) ` +
				`FROM table1 AS t1 ` +
				`RIGHT JOIN table2 AS t2 ` +
				`ON t1.column1 = t2.column1 ` +
				`GROUP BY column1 ` +
				`HAVING column1 > 10`,
		},
		{
			mgorm.Select(nil, "column1").
				From("table1 AS t1").
				FullJoin("table2 AS t2").
				On("t1.column1 = t2.column1").(*mgorm.Stmt),
			`SELECT column1 ` +
				`FROM table1 AS t1 ` +
				`FULL OUTER JOIN table2 AS t2 ` +
				`ON t1.column1 = t2.column1`,
		},
		{
			mgorm.Insert(nil, "table", "column1", "column2").
				Values(10, "str").(*mgorm.Stmt),
			`INSERT INTO table (column1, column2) ` +
				`VALUES (10, "str")`,
		},
		{
			mgorm.Update(nil, "table", "column1", "column2").
				Set(10, "str").
				Where("lhs = ?", 10).(*mgorm.Stmt),
			`UPDATE table ` +
				`SET column1 = 10, column2 = "str" ` +
				`WHERE lhs = 10`,
		},
		{
			mgorm.Delete(nil).
				From("table").(*mgorm.Stmt),
			`DELETE FROM table`,
		},
		{
			mgorm.Count(nil, "column", "c").
				From("table").(*mgorm.Stmt),
			`SELECT COUNT(column) AS c FROM table`,
		},
		{
			mgorm.Avg(nil, "column", "c").
				From("table").(*mgorm.Stmt),
			`SELECT AVG(column) AS c FROM table`,
		},
		{
			mgorm.Sum(nil, "column", "c").
				From("table").(*mgorm.Stmt),
			`SELECT SUM(column) AS c FROM table`,
		},
		{
			mgorm.Max(nil, "column", "c").
				From("table").(*mgorm.Stmt),
			`SELECT MAX(column) AS c FROM table`,
		},
		{
			mgorm.Min(nil, "column", "c").
				From("table").(*mgorm.Stmt),
			`SELECT MIN(column) AS c FROM table`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Stmt.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestStmt_ProcessQuerySQL_Fail(t *testing.T) {
	testCases := []struct {
		Cmd    syntax.Cmd
		Called []syntax.Expr
		Error  error
	}{
		{
			&cmd.Delete{},
			nil,
			errors.New("Command must be SELECT", errors.InvalidValueError),
		},
		{
			&cmd.Select{},
			[]syntax.Expr{
				&expr.From{},
				&expr.When{},
			},
			errors.New("Type expr.When is not supported", errors.InvalidTypeError),
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		stmt.ExportedSetCalled(testCase.Called)
		_, err := mgorm.StmtProcessQuerySQL(stmt)
		if err == nil {
			t.Errorf("Error was not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Type of error is invalid")
			continue
		}
		resultError := testCase.Error.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}

func TestStmt_ProcessCaseSQL_Fail(t *testing.T) {
	testCases := []struct {
		Called []syntax.Expr
		Error  error
	}{
		{
			[]syntax.Expr{
				&expr.When{},
				&expr.Where{},
			},
			errors.New("Type expr.Where is not supported", errors.InvalidTypeError),
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCalled(testCase.Called)
		_, err := mgorm.StmtProcessCaseSQL(stmt, false)
		if err == nil {
			t.Errorf("Error was not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Type of error is invalid")
			continue
		}
		resultError := testCase.Error.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}

func TestStmt_PrcessExecSQL_Fail(t *testing.T) {
	testCases := []struct {
		Cmd   syntax.Cmd
		Error error
	}{
		{
			&cmd.Select{},
			errors.New("Command must be INSERT, UPDATE or DELETE", errors.InvalidTypeError),
		},
	}

	for _, testCase := range testCases {
		stmt := new(mgorm.Stmt)
		stmt.ExportedSetCmd(testCase.Cmd)
		_, err := mgorm.StmtProcessExecSQL(stmt)
		if err == nil {
			t.Errorf("Error was not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Type of error is invalid")
			continue
		}
		resultError := testCase.Error.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}
