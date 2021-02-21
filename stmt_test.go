package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/stretchr/testify/assert"
)

func TestStmt_QuerySQLString(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.Stmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no = ?", 10001).(*mgorm.Stmt),
			`SELECT emp_no FROM employees WHERE emp_no = 10001`,
		},
		{
			mgorm.Select(nil, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10005).
				And("first_name = ? OR last_name = ?", "Georgi", "Bamford").(*mgorm.Stmt),
			`SELECT emp_no, first_name, last_name FROM employees ` +
				`WHERE emp_no <= 10005 ` +
				`AND (first_name = "Georgi" OR last_name = "Bamford")`,
		},
		{
			mgorm.Select(nil, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10002).
				Or("first_name = ? AND last_name = ?", "Saniya", "Kalloufi").(*mgorm.Stmt),
			`SELECT emp_no, first_name, last_name FROM employees ` +
				`WHERE emp_no <= 10002 ` +
				`OR (first_name = "Saniya" AND last_name = "Kalloufi")`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no IN ?",
					mgorm.Select(nil, "DISTINCT emp_no").
						From("salaries").
						Where("salary < ?", 60000)).(*mgorm.Stmt),
			`SELECT emp_no FROM employees ` +
				`WHERE emp_no IN ` +
				`(SELECT DISTINCT emp_no FROM salaries WHERE salary < 60000)`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no BETWEEN ? AND ?", 10002, 10004).(*mgorm.Stmt),
			`SELECT emp_no FROM employees WHERE emp_no BETWEEN 10002 AND 10004`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				Where("first_name LIKE ?", "S%").(*mgorm.Stmt),
			`SELECT first_name FROM employees WHERE first_name LIKE "S%"`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no IN (?, ?)", 10002, 10004).(*mgorm.Stmt),
			`SELECT emp_no FROM employees WHERE emp_no IN (10002, 10004)`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				OrderBy("first_name").(*mgorm.Stmt),
			`SELECT first_name FROM employees ORDER BY first_name`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				OrderBy("first_name", "last_name DESC").(*mgorm.Stmt),
			`SELECT first_name FROM employees ORDER BY first_name, last_name DESC`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Limit(5).(*mgorm.Stmt),
			`SELECT emp_no FROM employees LIMIT 5`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Limit(5).
				Offset(3).(*mgorm.Stmt),
			`SELECT emp_no FROM employees LIMIT 5 OFFSET 3`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "t.title").
				From("employees AS e").
				Join("titles AS t").
				On("e.emp_no = t.emp_no").
				Limit(5).(*mgorm.Stmt),
			`SELECT e.emp_no, t.title FROM employees AS e ` +
				`INNER JOIN titles AS t ` +
				`ON e.emp_no = t.emp_no ` +
				`LIMIT 5`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "t.title").
				From("employees AS e").
				LeftJoin("titles AS t").
				On("e.emp_no = t.emp_no").
				OrderBy("e.emp_no DESC", "title").
				Limit(5).(*mgorm.Stmt),
			`SELECT e.emp_no, t.title FROM employees AS e ` +
				`LEFT JOIN titles AS t ` +
				`ON e.emp_no = t.emp_no ` +
				`ORDER BY e.emp_no DESC, title ` +
				`LIMIT 5`,
		},
		{
			mgorm.Select(nil, "t.title", "e.emp_no").
				From("titles AS t").
				RightJoin("employees AS e").
				On("t.emp_no = e.emp_no").
				OrderBy("e.emp_no DESC", "title").
				Limit(5).(*mgorm.Stmt),
			`SELECT t.title, e.emp_no FROM titles AS t ` +
				`RIGHT JOIN employees AS e ` +
				`ON t.emp_no = e.emp_no ` +
				`ORDER BY e.emp_no DESC, title ` +
				`LIMIT 5`,
		},
		{
			mgorm.Select(nil, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").(*mgorm.Stmt),
			`SELECT title, COUNT(title) FROM titles GROUP BY title`,
		},
		{
			mgorm.Select(nil, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").
				Having("COUNT(title) != ?", 1).(*mgorm.Stmt),
			`SELECT title, COUNT(title) FROM titles ` +
				`GROUP BY title ` +
				`HAVING COUNT(title) != 1`,
		},
		{
			mgorm.Select(nil, "hire_date AS date").
				From("employees").
				Union(mgorm.Select(nil, "from_date AS date").
					From("salaries")).
				Limit(5).(*mgorm.Stmt),
			`SELECT hire_date AS date FROM employees ` +
				`UNION ` +
				`SELECT from_date AS date FROM salaries ` +
				`LIMIT 5`,
		},
		{
			mgorm.Select(nil, "from_date AS date").
				From("salaries").
				UnionAll(mgorm.Select(nil, "from_date AS date").
					From("titles")).
				Limit(5).(*mgorm.Stmt),
			`SELECT from_date AS date FROM salaries ` +
				`UNION ALL ` +
				`SELECT from_date AS date FROM titles ` +
				`LIMIT 5`,
		},
		{
			mgorm.Min(nil, "emp_no").From("employees").(*mgorm.Stmt),
			`SELECT MIN(emp_no) FROM employees`,
		},
		{
			mgorm.Max(nil, "emp_no").From("employees").(*mgorm.Stmt),
			`SELECT MAX(emp_no) FROM employees`,
		},
		{
			mgorm.Count(nil, "emp_no").From("salaries").(*mgorm.Stmt),
			`SELECT COUNT(emp_no) FROM salaries`,
		},
		{
			mgorm.Sum(nil, "salary").From("salaries").(*mgorm.Stmt),
			`SELECT SUM(salary) FROM salaries`,
		},
		{
			mgorm.Avg(nil, "salary").From("salaries").(*mgorm.Stmt),
			`SELECT AVG(salary) FROM salaries`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestStmt_ExecSQLString(t *testing.T) {
	type Model struct {
		ID   int
		Name string `mgorm:"first_name"`
	}

	model1 := Model{ID: 10000, Name: "Taro"}
	model2 := []Model{{ID: 10000, Name: "Taro"}, {ID: 10001, Name: "Hanako"}}
	model3 := []int{10000, 10001}

	testCases := []struct {
		Stmt     *mgorm.Stmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "sample", "name AS first_name", "id").Model(&model1).(*mgorm.Stmt),
			`INSERT INTO sample (name AS first_name, id) VALUES ("Taro", 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name AS first_name", "id").Model(&model2).(*mgorm.Stmt),
			`INSERT INTO sample (name AS first_name, id) VALUES ("Taro", 10000), ("Hanako", 10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "id").Model(&model3).(*mgorm.Stmt),
			`INSERT INTO sample (id) VALUES (10000), (10001)`,
		},
		{
			mgorm.Update(nil, "sample", "id", "first_name").Model(&model1).Where("id = ?", 20000).(*mgorm.Stmt),
			`UPDATE sample SET id = 10000, first_name = "Taro" WHERE id = 20000`,
		},
		{
			mgorm.Update(nil, "sample", "id").Model(10000).Where("id = ?", 20000).(*mgorm.Stmt),
			`UPDATE sample SET id = 10000 WHERE id = 20000`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestStmt_ProcessQuerySQL_Fail(t *testing.T) {
	{
		expectedErr := errors.New("Command must be SELECT", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Update(nil, "column1", "column2").Set(10, "str").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.StmtProcessQuerySQL(s)

		// Validate error.
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		expectedErr := errors.New(
			`Type clause.Values is not supported for SELECT`, errors.InvalidTypeError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Select(nil, "").(*mgorm.Stmt).Values("").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.StmtProcessQuerySQL(s)

		// Validate error.
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestStmt_ProcessExecSQL_Fail(t *testing.T) {
	{
		expectedErr := errors.New(
			"Command must be INSERT, UPDATE or DELETE", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Select(nil, "").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.StmtProcessExecSQL(s)

		// Validate error.
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		expectedErr := errors.New(
			"Type clause.Join is not supported for INSERT", errors.InvalidTypeError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Insert(nil, "").(*mgorm.Stmt).Join("").(*mgorm.Stmt)

		// Actual process.
		_, err := mgorm.StmtProcessExecSQL(s)

		// Validate error.
		if err == nil {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}

func TestStmt_Set_Fail(t *testing.T) {
	{
		expectedErr := errors.New("Command is nil", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := new(mgorm.Stmt)

		// Actual process.
		s.Set("")

		// Validate error.
		errs := s.ExportedGetErrors()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
	{
		expectedErr := errors.New(
			"SET clause can be used with UPDATE command", errors.InvalidValueError).(*errors.Error)

		// Prepare for test.
		s := mgorm.Select(nil, "").(*mgorm.Stmt).Set("").(*mgorm.Stmt)

		// Actual process.
		s.Set("")

		// Validate error.
		errs := s.ExportedGetErrors()
		if len(errs) == 0 {
			t.Errorf("Error was not occurred")
			return
		}
		actualErr, ok := errs[0].(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			return
		}
		if !actualErr.Is(expectedErr) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
		}
	}
}
