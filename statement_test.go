package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_BuildSQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Delete(nil).(*mgorm.DeleteStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.DeleteStmtBuildSQL(s, &sql)
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

func TestInsertStmt_String(t *testing.T) {
	type Model struct {
		ID        int
		FirstName string `mgorm:"name"`
	}

	model1 := Model{ID: 10000, FirstName: "Taro"}
	model2 := []Model{{ID: 10000, FirstName: "Taro"}, {ID: 10001, FirstName: "Jiro"}}
	model3 := []int{10000, 10001}
	model4 := map[string]interface{}{
		"id":   10000,
		"name": "Taro",
	}

	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "sample", "id", "name").Values(10000, "Taro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro'), (10001, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").
				Values(10000, "Taro").
				Values(10001, "Jiro").
				Values(10002, "Saburo").(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro'), (10001, 'Jiro'), (10002, 'Saburo')`,
		},
		// Test for (*InsertStmt).Model
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro'), (10001, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "sample", "id").Model(&model3).(*mgorm.InsertStmt),
			`INSERT INTO sample (id) VALUES (10000), (10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "id", "name").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (id, name) VALUES (10000, 'Taro')`,
		},
		// Test for mapping.
		{
			mgorm.Insert(nil, "sample", "first_name AS name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (first_name AS name, id) VALUES ('Taro', 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model1).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ('Taro', 10000)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model2).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ('Taro', 10000), ('Jiro', 10001)`,
		},
		{
			mgorm.Insert(nil, "sample", "name", "id").Model(&model4).(*mgorm.InsertStmt),
			`INSERT INTO sample (name, id) VALUES ('Taro', 10000)`,
		},
		// Test for INSERT INTO ... SELECT statement.
		{
			mgorm.Insert(nil, "person").
				Select(mgorm.Select(nil, "id", "name").
					From("country_code").
					Where("name = ?", "Japan"),
				).(*mgorm.InsertStmt),
			`INSERT INTO person SELECT id, name FROM country_code WHERE name = 'Japan'`,
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

func TestInsertStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").(*mgorm.InsertStmt)
				s.ExportedSetCalled(&clause.Set{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
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

func TestInsertStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				s := mgorm.Insert(nil, "", "").Model(1000).(*mgorm.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := mgorm.Insert(nil, "table", "column").Model(&model).(*mgorm.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.InsertStmtBuildSQL(s, &sql)
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

func TestSelectStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no = ?", 10001).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees WHERE emp_no = 10001`,
		},
		{
			mgorm.Select(nil, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10005).
				And("first_name = ? OR last_name = ?", "Georgi", "Bamford").(*mgorm.SelectStmt),
			`SELECT emp_no, first_name, last_name FROM employees ` +
				`WHERE emp_no <= 10005 ` +
				`AND (first_name = 'Georgi' OR last_name = 'Bamford')`,
		},
		{
			mgorm.Select(nil, "emp_no", "first_name", "last_name").
				From("employees").
				Where("emp_no <= ?", 10002).
				Or("first_name = ? AND last_name = ?", "Saniya", "Kalloufi").(*mgorm.SelectStmt),
			`SELECT emp_no, first_name, last_name FROM employees ` +
				`WHERE emp_no <= 10002 ` +
				`OR (first_name = 'Saniya' AND last_name = 'Kalloufi')`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no IN ?",
					mgorm.Select(nil, "DISTINCT emp_no").
						From("salaries").
						Where("salary < ?", 60000)).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees ` +
				`WHERE emp_no IN ` +
				`(SELECT DISTINCT emp_no FROM salaries WHERE salary < 60000)`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no BETWEEN ? AND ?", 10002, 10004).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees WHERE emp_no BETWEEN 10002 AND 10004`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				Where("first_name LIKE ?", "S%").(*mgorm.SelectStmt),
			`SELECT first_name FROM employees WHERE first_name LIKE 'S%'`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Where("emp_no IN (?, ?)", 10002, 10004).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees WHERE emp_no IN (10002, 10004)`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				OrderBy("first_name").(*mgorm.SelectStmt),
			`SELECT first_name FROM employees ORDER BY first_name`,
		},
		{
			mgorm.Select(nil, "first_name").
				From("employees").
				OrderBy("first_name", "last_name DESC").(*mgorm.SelectStmt),
			`SELECT first_name FROM employees ORDER BY first_name, last_name DESC`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Limit(5).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees LIMIT 5`,
		},
		{
			mgorm.Select(nil, "emp_no").
				From("employees").
				Limit(5).
				Offset(3).(*mgorm.SelectStmt),
			`SELECT emp_no FROM employees LIMIT 5 OFFSET 3`,
		},
		{
			mgorm.Select(nil, "e.emp_no", "t.title").
				From("employees AS e").
				Join("titles AS t").
				On("e.emp_no = t.emp_no").
				Limit(5).(*mgorm.SelectStmt),
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
				Limit(5).(*mgorm.SelectStmt),
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
				Limit(5).(*mgorm.SelectStmt),
			`SELECT t.title, e.emp_no FROM titles AS t ` +
				`RIGHT JOIN employees AS e ` +
				`ON t.emp_no = e.emp_no ` +
				`ORDER BY e.emp_no DESC, title ` +
				`LIMIT 5`,
		},
		{
			mgorm.Select(nil).
				From("employees AS e").
				Join("titles AS t").On("t.emp_no = e.emp_no").
				LeftJoin("departments AS d").On("d.dept_no = e.dept_no").(*mgorm.SelectStmt),
			`SELECT * FROM employees AS e ` +
				`INNER JOIN titles AS t ON t.emp_no = e.emp_no ` +
				`LEFT JOIN departments AS d ON d.dept_no = e.dept_no`,
		},
		{
			mgorm.Select(nil, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").(*mgorm.SelectStmt),
			`SELECT title, COUNT(title) FROM titles GROUP BY title`,
		},
		{
			mgorm.Select(nil, "title", "COUNT(title)").
				From("titles").
				GroupBy("title").
				Having("COUNT(title) != ?", 1).(*mgorm.SelectStmt),
			`SELECT title, COUNT(title) FROM titles ` +
				`GROUP BY title ` +
				`HAVING COUNT(title) != 1`,
		},
		{
			mgorm.Select(nil, "hire_date AS date").
				From("employees").
				Union(mgorm.Select(nil, "from_date AS date").
					From("salaries")).
				Limit(5).(*mgorm.SelectStmt),
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
				Limit(5).(*mgorm.SelectStmt),
			`SELECT from_date AS date FROM salaries ` +
				`UNION ALL ` +
				`SELECT from_date AS date FROM titles ` +
				`LIMIT 5`,
		},
		{
			mgorm.Min(nil, "emp_no").From("employees").(*mgorm.SelectStmt),
			`SELECT MIN(emp_no) FROM employees`,
		},
		{
			mgorm.Max(nil, "emp_no").From("employees").(*mgorm.SelectStmt),
			`SELECT MAX(emp_no) FROM employees`,
		},
		{
			mgorm.Count(nil, "emp_no").From("salaries").(*mgorm.SelectStmt),
			`SELECT COUNT(emp_no) FROM salaries`,
		},
		{
			mgorm.Sum(nil, "salary").From("salaries").(*mgorm.SelectStmt),
			`SELECT SUM(salary) FROM salaries`,
		},
		{
			mgorm.Avg(nil, "salary").From("salaries").(*mgorm.SelectStmt),
			`SELECT AVG(salary) FROM salaries`,
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

func TestStmt_BuildQuerySQL_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Select(nil, "").(*mgorm.SelectStmt)
				s.ExportedSetCalled(&clause.Values{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.SelectStmtBuildSQL(s, &sql)
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

func TestUpdateStmt_String(t *testing.T) {
	type Model struct {
		ID   int
		Name string `mgorm:"first_name"`
	}

	model1 := Model{ID: 10000, Name: "Taro"}
	model2 := map[string]interface{}{"id": 10000, "first_name": "Taro"}

	testCases := []struct {
		Stmt     *mgorm.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "sample").
				Set("id", 10000).
				Set("first_name", "Taro").
				Where("id = ?", 20000).
				And("first_name = ? OR first_name = ?", "Jiro", "Hanako").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000 ` +
				`AND (first_name = 'Jiro' OR first_name = 'Hanako')`,
		},
		{
			mgorm.Update(nil, "sample").
				Set("id", 10000).
				Set("first_name", "Taro").
				Where("id = ?", 20000).
				Or("first_name = ? AND last_name = ?", "Jiro", "Sato").(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000 ` +
				`OR (first_name = 'Jiro' AND last_name = 'Sato')`,
		},
		{
			mgorm.Update(nil, "sample").
				Model(&model1, "id", "first_name").
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000`,
		},
		{
			mgorm.Update(nil, "sample").
				Model(&model2, "id", "first_name").
				Where("id = ?", 20000).(*mgorm.UpdateStmt),
			`UPDATE sample SET id = 10000, first_name = 'Taro' ` +
				`WHERE id = 20000`,
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

func TestUpdateStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := mgorm.Update(nil, "table").(*mgorm.UpdateStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
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

func TestUpdateStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := map[string]interface{}{
					"id":   1000,
					"name": "Taro",
				}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*mgorm.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			mgorm.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := []int{1000}
				s := mgorm.Update(nil, "sample").Model(&model, "id", "first_name").(*mgorm.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := mgorm.UpdateStmtBuildSQL(s, &sql)
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
