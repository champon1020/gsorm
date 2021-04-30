package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").(*statement.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees").(*statement.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees AS e").(*statement.DeleteStmt),
			`DELETE FROM employees AS e`,
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

func TestDeleteStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = 1001").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("first_name = ?", "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE first_name = 'Taro'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
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

func TestDeleteStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestDeleteStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestInsertStmt_Values(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").
				Values(1002, "Jiro").(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
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

func TestInsertStmt_Select(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "dept_manager").Select(mgorm.Select(nil).From("dept_emp")).(*statement.InsertStmt),
			`INSERT INTO dept_manager SELECT * FROM dept_emp`,
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

func TestInsertStmt_Model(t *testing.T) {
	type Employee struct {
		ID        int `mgorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	structSlice := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}
	mapSlice := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
	}
	varSlice := []string{"Taro", "Jiro"}

	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structModel).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structSlice).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapModel).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapSlice).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "employees", "first_name").Model(&varSlice).(*statement.InsertStmt),
			`INSERT INTO employees (first_name) VALUES ('Taro'), ('Jiro')`,
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

func TestUpdateStmt_Set(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako'`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Set("last_name", "Suzuki").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako', last_name = 'Suzuki'`,
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

func TestUpdateStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = 1001").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name = ?", "Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name = 'Taro'`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name LIKE '%Taro'`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
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

func TestUpdateStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002) ` +
				`AND (emp_no = 1003)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestUpdateStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002) ` +
				`OR (emp_no = 1003)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestUpdateStmt_Model(t *testing.T) {
	type Employee struct {
		ID        int `mgorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			mgorm.Update(nil, "employees").Model(&structModel, "emp_no", "first_name").(*statement.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
		},
		{
			mgorm.Update(nil, "employees").Model(&mapModel, "emp_no", "first_name").(*statement.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
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
