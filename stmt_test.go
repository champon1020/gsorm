package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/champon1020/mgorm/syntax/cmd"
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
				OrderBy("t1.column1").
				UnionAll(mgorm.Select(nil, "column1", "column2 AS c2").
					From("table3").
					Sub()).(*mgorm.Stmt),
			`SELECT column1, column2 AS c2 ` +
				`FROM table1 AS t1 ` +
				`LEFT JOIN table2 AS t2 ` +
				`ON t1.column1 = t2.column1 ` +
				`WHERE lhs1 > 10 ` +
				`OR (lhs2 = "str" AND lhs3 = true) ` +
				`ORDER BY t1.column1 ` +
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
			mgorm.Select(nil, mgorm.When("column1 < ?", 10).
				Then("column1").
				When("column1 > ?", 10).
				Then("column2").
				Else("column3").CaseColumn()).
				From("table1").(*mgorm.Stmt),
			`SELECT CASE WHEN column1 < 10 THEN column1 ` +
				`WHEN column1 > 10 THEN column2 ` +
				`ELSE column3 END ` +
				`FROM table1`,
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
		Called []syntax.Clause
		Error  error
	}{
		{
			&cmd.Delete{},
			nil,
			errors.New("Command must be SELECT", errors.InvalidValueError),
		},
		{
			&cmd.Select{},
			[]syntax.Clause{
				&clause.From{},
				&clause.When{},
			},
			errors.New("Type clause.When is not supported", errors.InvalidTypeError),
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
		Called []syntax.Clause
		Error  error
	}{
		{
			[]syntax.Clause{
				&clause.When{},
				&clause.Where{},
			},
			errors.New("Type clause.Where is not supported", errors.InvalidTypeError),
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
			errors.New("Command must be INSERT, UPDATE or DELETE", errors.InvalidValueError),
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

func TestStmt_Set_Fail(t *testing.T) {
	testCases := []struct {
		Cmd    syntax.Cmd
		Values []interface{}
		Error  error
	}{
		{
			nil,
			[]interface{}{10},
			errors.New("Command is nil", errors.InvalidValueError),
		},
		{
			&cmd.Select{},
			[]interface{}{10},
			errors.New("SET clause can be used with UPDATE command", errors.InvalidValueError),
		},
		{
			&cmd.Update{Columns: []string{"lhs1", "lhs2"}},
			[]interface{}{10},
			errors.New("Length is different between lhs and rhs", errors.InvalidValueError),
		},
	}

	for _, testCase := range testCases {
		s := new(mgorm.Stmt)
		s.ExportedSetCmd(testCase.Cmd)
		resultStmt := s.Set(testCase.Values...).(*mgorm.Stmt)
		errs := resultStmt.ExportedGetErrors()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			continue
		}
		actualError, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
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
