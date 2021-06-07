package gsorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestStatement_QueryWithMock(t *testing.T) {
	type Employee struct {
		EmpNo     int
		FirstName string
	}
	model := []Employee{}

	mock := gsorm.OpenMock()
	expectedReturn := []Employee{
		{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"},
	}
	mock.ExpectWithReturn(gsorm.Select(mock, "emp_no", "first_name").From("employees"), expectedReturn)

	err := gsorm.Select(mock, "emp_no", "first_name").From("employees").Query(&model)
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if diff := cmp.Diff(expectedReturn, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestStatement_ExecWithMock(t *testing.T) {
	mock := gsorm.OpenMock()
	mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))

	err := gsorm.Insert(mock, "employees", "emp_no", "first_name").Values(1001, "Taro").Exec()
	if err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}
}

func TestDeleteStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).
				RawClause("RAW").
				From("employees").(*gsorm.DeleteStmt),
			`DELETE RAW FROM employees`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				RawClause("RAW").
				Where("emp_no = ?", 10000).(*gsorm.DeleteStmt),
			`DELETE FROM employees RAW WHERE emp_no = 10000`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				And("first_name = ?", "Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW AND (first_name = 'Taro')`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				RawClause("RAW").
				Or("emp_no = ?", 20000).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 RAW OR (emp_no = 20000)`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				And("first_name = ?", "Taro").
				RawClause("RAW").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 AND (first_name = 'Taro') RAW`,
		},
		{
			gsorm.Delete(nil).
				From("employees").
				Where("emp_no = ?", 10000).
				Or("emp_no = ?", 20000).
				RawClause("RAW").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 10000 OR (emp_no = 20000) RAW`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestDeleteStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").(*gsorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.Delete(nil).From("employees").(*gsorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.Delete(nil).From("employees AS e").(*gsorm.DeleteStmt),
			`DELETE FROM employees AS e`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = 1001").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("first_name = ?", "Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE first_name = 'Taro'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.DeleteStmt
		Expected string
	}{
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestInsertStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "table").
				RawClause("RAW").
				Values("value").(*gsorm.InsertStmt),
			`INSERT INTO table RAW VALUES ('value')`,
		},
		{
			gsorm.Insert(nil, "table").
				Values("value1").
				RawClause("RAW").
				Values("value2").(*gsorm.InsertStmt),
			`INSERT INTO table VALUES ('value1') RAW, ('value2')`,
		},
		{
			gsorm.Insert(nil, "table").
				Values("value").
				RawClause("RAW").(*gsorm.InsertStmt),
			`INSERT INTO table VALUES ('value') RAW`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "employees").
				Values(1001, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").(*gsorm.InsertStmt),
			`INSERT INTO employees VALUES (1001, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").
				Values(1002, "Jiro").(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "dept_manager").Select(gsorm.Select(nil).From("dept_emp")).(*gsorm.InsertStmt),
			`INSERT INTO dept_manager SELECT * FROM dept_emp`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		ID        int `gsorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	structSlice := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}
	mapSlice := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
	}

	testCases := []struct {
		Stmt     *gsorm.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structModel).(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structSlice).(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapModel).(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapSlice).(*gsorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).
				RawClause("RAW").
				From("table").(*gsorm.SelectStmt),
			`SELECT * RAW FROM table`,
		},
		{
			gsorm.Select(nil).
				From("table").
				RawClause("RAW").
				Join("table2").On("table.column = table2.column").(*gsorm.SelectStmt),
			`SELECT * FROM table RAW ` +
				`INNER JOIN table2 ON table.column = table2.column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				RawClause("RAW").
				LeftJoin("table2").On("table.column = table2.column").(*gsorm.SelectStmt),
			`SELECT * FROM table RAW ` +
				`LEFT JOIN table2 ON table.column = table2.column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				RawClause("RAW").
				RightJoin("table2").On("table.column = table2.column").(*gsorm.SelectStmt),
			`SELECT * FROM table RAW ` +
				`RIGHT JOIN table2 ON table.column = table2.column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Join("table2").RawClause("RAW").On("table.column = table2.column").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 RAW ON table.column = table2.column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Join("table2").On("table.column = table2.column").
				RawClause("RAW").
				Where("id = ?", 10).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`INNER JOIN table2 ON table.column = table2.column ` +
				`RAW WHERE id = 10`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				And("name = ?", "Taro").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW AND (name = 'Taro')`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				RawClause("RAW").
				Or("name = ?", "Taro").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 RAW OR (name = 'Taro')`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				And("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 AND (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Where("id = ?", 10).
				Or("name = ?", "Taro").
				RawClause("RAW").
				GroupBy("column").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`WHERE id = 10 OR (name = 'Taro') ` +
				`RAW GROUP BY column`,
		},
		{
			gsorm.Select(nil).
				From("table").
				GroupBy("column").
				RawClause("RAW").
				Having("SUM(id) = ?", 10).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`GROUP BY column ` +
				`RAW HAVING SUM(id) = 10`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				Union(gsorm.Select(nil).From("table2")).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION (SELECT * FROM table2)`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Having("SUM(id) = ?", 10).
				RawClause("RAW").
				UnionAll(gsorm.Select(nil).From("table2")).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`HAVING SUM(id) = 10 ` +
				`RAW UNION ALL (SELECT * FROM table2)`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Union(gsorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`UNION (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			gsorm.Select(nil).
				From("table").
				UnionAll(gsorm.Select(nil).From("table2")).
				RawClause("RAW").
				OrderBy("id").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`UNION ALL (SELECT * FROM table2) ` +
				`RAW ORDER BY id`,
		},
		{
			gsorm.Select(nil).
				From("table").
				OrderBy("id").
				RawClause("RAW").
				Limit(10).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`ORDER BY id ` +
				`RAW LIMIT 10`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Limit(10).
				RawClause("RAW").
				Offset(5).(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 RAW OFFSET 5`,
		},
		{
			gsorm.Select(nil).
				From("table").
				Limit(10).
				Offset(5).
				RawClause("RAW").(*gsorm.SelectStmt),
			`SELECT * FROM table ` +
				`LIMIT 10 OFFSET 5 RAW`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "emp_no").From("employees").(*gsorm.SelectStmt),
			`SELECT emp_no FROM employees`,
		},
		{
			gsorm.Select(nil, "emp_no").From("employees AS e").(*gsorm.SelectStmt),
			`SELECT emp_no FROM employees AS e`,
		},
		{
			gsorm.Select(nil, "emp_no").From("employees as e").(*gsorm.SelectStmt),
			`SELECT emp_no FROM employees AS e`,
		},
		{
			gsorm.Select(nil, "emp_no", "dept_no").From("employees", "departments").(*gsorm.SelectStmt),
			`SELECT emp_no, dept_no FROM employees, departments`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				Join("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				Join("dept_manager AS d").On("e.emp_no = d.emp_no").
				LeftJoin("salaries AS s").On("e.emp_no = s.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`LEFT JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				LeftJoin("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`LEFT JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				LeftJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
				RightJoin("salaries AS s").On("e.emp_no = s.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`LEFT JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`RIGHT JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				RightJoin("dept_manager AS d").
				On("e.emp_no = d.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no`,
		},
		{
			gsorm.Select(nil, "e.emp_no", "d.dept_no").
				From("employees AS e").
				RightJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
				Join("salaries AS s").On("e.emp_no = s.emp_no").(*gsorm.SelectStmt),
			`SELECT e.emp_no, d.dept_no FROM employees AS e ` +
				`RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no ` +
				`INNER JOIN salaries AS s ON e.emp_no = s.emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = 1001").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("first_name = ?", "Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE first_name = 'Taro'`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Select(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.SelectStmt),
			`SELECT * FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries GROUP BY emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) > 130000").(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) > 130000`,
		},
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) > ?", 130000).(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) > 130000`,
		},
		{
			gsorm.Select(nil).From("employees").
				Having("first_name = ?", "Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees HAVING first_name = 'Taro'`,
		},
		{
			gsorm.Select(nil).From("employees").
				Having("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.SelectStmt),
			`SELECT * FROM employees HAVING birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Select(nil).From("employees").
				Having("first_name LIKE ?", "%Taro").(*gsorm.SelectStmt),
			`SELECT * FROM employees HAVING first_name LIKE '%Taro'`,
		},
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) BETWEEN ? AND ?", 100000, 130000).(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) BETWEEN 100000 AND 130000`,
		},
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) IN (?)", []int{100000, 130000}).(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) IN (100000, 130000)`,
		},
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				GroupBy("emp_no").
				Having("AVG(salary) IN (?)", [2]int{100000, 130000}).(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`GROUP BY emp_no ` +
				`HAVING AVG(salary) IN (100000, 130000)`,
		},
		{
			gsorm.Select(nil, "emp_no", "AVG(salary)").From("salaries").
				Having("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.SelectStmt),
			`SELECT emp_no, AVG(salary) FROM salaries ` +
				`HAVING emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "emp_no", "dept_no").From("dept_manager").
				Union(gsorm.Select(nil, "emp_no", "dept_no").From("dept_emp")).(*gsorm.SelectStmt),
			`SELECT emp_no, dept_no FROM dept_manager ` +
				`UNION (SELECT emp_no, dept_no FROM dept_emp)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil, "emp_no", "dept_no").From("dept_manager").
				UnionAll(gsorm.Select(nil, "emp_no", "dept_no").From("dept_emp")).(*gsorm.SelectStmt),
			`SELECT emp_no, dept_no FROM dept_manager ` +
				`UNION ALL (SELECT emp_no, dept_no FROM dept_emp)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).From("employees").
				OrderBy("birth_date").(*gsorm.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date`,
		},
		{
			gsorm.Select(nil).From("employees").
				OrderBy("birth_date DESC").(*gsorm.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date DESC`,
		},
		{
			gsorm.Select(nil).From("employees").
				OrderBy("birth_date desc").(*gsorm.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date desc`,
		},
		{
			gsorm.Select(nil).From("employees").
				OrderBy("birth_date").
				OrderBy("hire_date DESC").(*gsorm.SelectStmt),
			`SELECT * FROM employees ORDER BY birth_date ORDER BY hire_date DESC`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.SelectStmt
		Expected string
	}{
		{
			gsorm.Select(nil).From("employees").Limit(10).(*gsorm.SelectStmt),
			`SELECT * FROM employees LIMIT 10`,
		},
		{
			gsorm.Select(nil).From("employees").Limit(10).Offset(5).(*gsorm.SelectStmt),
			`SELECT * FROM employees LIMIT 10 OFFSET 5`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestUpdateStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "table").
				RawClause("RAW").
				Set("column", "value").(*gsorm.UpdateStmt),
			`UPDATE table RAW SET column = 'value'`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Set("column2", "value2").(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' RAW, column2 = 'value2'`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Where("column = ?", 10).(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`RAW WHERE column = 10`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				And("column = ?", 100).(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW AND (column = 100)`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				Or("column = ?", 100).(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW OR (column = 100)`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				And("column = ?", 100).
				RawClause("RAW").(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`AND (column = 100) RAW`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				Or("column = ?", 100).
				RawClause("RAW").(*gsorm.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`OR (column = 100) RAW`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Set("last_name", "Suzuki").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako', last_name = 'Suzuki'`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = 1001").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name = ?", "Taro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name = 'Taro'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name LIKE ?", "%Taro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name LIKE '%Taro'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", []int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002) ` +
				`AND (emp_no = 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name = ? OR first_name = ?", "Taro", "Jiro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002) ` +
				`OR (emp_no = 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
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
		ID        int `gsorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	testCases := []struct {
		Stmt     *gsorm.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").Model(&structModel, "emp_no", "first_name").(*gsorm.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
		},
		{
			gsorm.Update(nil, "employees").Model(&mapModel, "emp_no", "first_name").(*gsorm.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestRawStmt_String(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.ExportedRawStmt
		Expected string
	}{
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees").(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no = ?", 1001).(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees WHERE emp_no = 1001`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE first_name = ?", "Taro").(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees WHERE first_name = 'Taro'`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE birth_date = ?",
				time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)", []int{1001, 1002}).(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.RawStmt(nil, "SELECT * FROM employees WHERE emp_no IN (?)",
				gsorm.Select(nil, "emp_no").From("dept_manager")).(*gsorm.ExportedRawStmt),
			`SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
		{
			gsorm.RawStmt(nil, "DELETE FROM employees").(*gsorm.ExportedRawStmt),
			`DELETE FROM employees`,
		},
		{
			gsorm.RawStmt(nil, "ALTER TABLE employees DROP PRIMARY KEY PK_emp_no").(*gsorm.ExportedRawStmt),
			`ALTER TABLE employees DROP PRIMARY KEY PK_emp_no`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.SQL()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestRawStmt_QueryWithMock(t *testing.T) {
	type Employee struct {
		EmpNo     int
		FirstName string
	}
	model := []Employee{}
	expectedReturn := []Employee{{1001, "Taro"}, {1002, "Jiro"}}

	mock := gsorm.OpenMock()
	mock.ExpectWithReturn(gsorm.RawStmt(nil, "SELECT emp_no, first_name FROM employees"), expectedReturn)

	if err := gsorm.RawStmt(mock, "SELECT emp_no, first_name FROM employees").Query(&model); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if diff := cmp.Diff(expectedReturn, model); diff != "" {
		t.Errorf("Differs: (-want +got)\n%s", diff)
	}
}

func TestRawStmt_ExecWithMock(t *testing.T) {
	mock := gsorm.OpenMock()
	mock.Expect(gsorm.RawStmt(mock, `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`))

	if err := gsorm.RawStmt(mock, `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`).Exec(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}

	if err := mock.Complete(); err != nil {
		t.Errorf("Error was occurred: %v", err)
	}
}
