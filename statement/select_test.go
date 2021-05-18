package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/statement"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestSelectStmt_BuildQuerySQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Select(nil, "").(*statement.SelectStmt)
				s.ExportedSetCalled(&clause.Values{})

				// Actual build.
				var sql internal.SQL
				err := statement.SelectStmtBuildSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Build()
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestSelectStmt_CompareStmts_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.SelectStmt
		ActualStmt    *statement.SelectStmt
		ExpectedError failure.StringCode
	}{
		{
			mgorm.Select(nil, "column1").From("table").(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			statement.ErrInvalidValue,
		},
		{
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 10).(*statement.SelectStmt),
			mgorm.Select(nil, "column1").From("table").Where("column1 = ?", 100).(*statement.SelectStmt),
			statement.ErrInvalidValue,
		},
	}

	for _, testCase := range testCases {
		err := testCase.ExpectedStmt.CompareWith(testCase.ActualStmt)

		// Validate if the expected error was occurred.
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestSelectStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).
				RawClause("RAW").
				From("table").(*statement.SelectStmt),
			`SELECT * RAW FROM table`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				Join("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`INNER JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				LeftJoin("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`LEFT JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				RawClause("RAW").
				RightJoin("table2").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table RAW ` +
				`RIGHT JOIN table2 ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Join("table2").RawClause("RAW").On("table.column = table2.column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 RAW ON table.column = table2.column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Join("table2").On("table.column = table2.column").
				RawClause("RAW").
				Where("id = ?", 10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 ON table.column = table2.column ` +
				`RAW WHERE id = 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				And("name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW AND (name = 'Taro')`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				Or("name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW OR (name = 'Taro')`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				And("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 AND (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				Or("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 OR (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			mgorm.Select(nil).
				From("table").
				GroupBy("column").
				RawClause("RAW").
				Having("SUM(id) = ?", 10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`GROUP BY column ` +
				`RAW HAVING SUM(id) = 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				Union(mgorm.Select(nil).From("table2")).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION (SELECT * FROM table2)`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				UnionAll(mgorm.Select(nil).From("table2")).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION ALL (SELECT * FROM table2)`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Union(mgorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`UNION (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			mgorm.Select(nil).
				From("table").
				UnionAll(mgorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`UNION ALL (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			mgorm.Select(nil).
				From("table").
				OrderBy("id").
				RawClause("RAW").
				Limit(10).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`ORDER BY id ` +
				`RAW LIMIT 10`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Limit(10).
				RawClause("RAW").
				Offset(5).(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 RAW OFFSET 5`,
		},
		{
			mgorm.Select(nil).
				From("table").
				Limit(10).
				Offset(5).
				RawClause("RAW").(*statement.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 OFFSET 5 RAW`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no").From("employees").(*statement.SelectStmt),
			`SELECT emp_no FROM employees`,
		},
		{
			mgorm.Select(nil, "emp_no").From("employees AS e").(*statement.SelectStmt),
			`SELECT emp_no FROM employees AS e`,
		},
		{
			mgorm.Select(nil, "emp_no").From("employees as e").(*statement.SelectStmt),
			`SELECT emp_no FROM employees AS e`,
		},
		{
			mgorm.Select(nil, "emp_no", "dept_no").From("employees", "departments").(*statement.SelectStmt),
			`SELECT emp_no, dept_no FROM employees, departments`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Join(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				Join("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				Join("dept_manager AS d").On("e.emp_no = d.emp_no").
				LeftJoin("salaries AS s").On("e.emp_no = s.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`LEFT JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_LeftJoin(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				LeftJoin("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`LEFT JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				LeftJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
				RightJoin("salaries AS s").On("e.emp_no = s.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`LEFT JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`RIGHT JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_RightJoin(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				RightJoin("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				RightJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
				Join("salaries AS s").On("e.emp_no = s.emp_no").(*statement.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`INNER JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = 1001").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("first_name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE first_name = 'Taro'`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_GroupBy(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries GROUP BY emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Having(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) > 130000").(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) > 130000`,
		},
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) > ?", 130000).(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) > 130000`,
		},
		{
			mgorm.Select(nil).From("employees").
				Having("first_name = ?", "Taro").(*statement.SelectStmt),
			`SELECT * FROM employees HAVING first_name = 'Taro'`,
		},
		{
			mgorm.Select(nil).From("employees").
				Having("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.SelectStmt),
			`SELECT * FROM employees HAVING birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.Select(nil).From("employees").
				Having("first_name LIKE ?", "%Taro").(*statement.SelectStmt),
			`SELECT * FROM employees HAVING first_name LIKE '%Taro'`,
		},
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) BETWEEN ? AND ?", 100000, 130000).(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) BETWEEN 100000 AND 130000`,
		},
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) IN (?)", []int{100000, 130000}).(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) IN (100000, 130000)`,
		},
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) IN (?)", [2]int{100000, 130000}).(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) IN (100000, 130000)`,
		},
		{
			mgorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				Having("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`HAVING emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Union(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no", "dept_no").From("dept_manager").
				Union(mgorm.Select(nil, "emp_no", "dept_no").From("dept_emp")).(*statement.SelectStmt),
			`SELECT emp_no, dept_no FROM dept_manager ` +
				`UNION (SELECT emp_no, dept_no FROM dept_emp)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_UnionAll(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no", "dept_no").From("dept_manager").
				UnionAll(mgorm.Select(nil, "emp_no", "dept_no").From("dept_emp")).(*statement.SelectStmt),
			`SELECT emp_no, dept_no FROM dept_manager ` +
				`UNION ALL (SELECT emp_no, dept_no FROM dept_emp)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_OrderBy(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).From("employees").
				OrderBy("birth_date").(*statement.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date`,
		},
		{
			mgorm.Select(nil).From("employees").
				OrderBy("birth_date DESC").(*statement.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date DESC`,
		},
		{
			mgorm.Select(nil).From("employees").
				OrderBy("birth_date desc").(*statement.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date desc`,
		},
		{
			mgorm.Select(nil).From("employees").
				OrderBy("birth_date").
				OrderBy("hire_date DESC").(*statement.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date ORDER BY hire_date DESC`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Limit(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil).From("employees").Limit(10).(*statement.SelectStmt),
			`SELECT * FROM employees LIMIT 10`,
		},
		{
			mgorm.Select(nil).From("employees").Limit(10).Offset(5).(*statement.SelectStmt),
			`SELECT * FROM employees LIMIT 10 OFFSET 5`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
